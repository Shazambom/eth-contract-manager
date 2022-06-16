package storage

import (
	"context"
	pb "contract-service/proto"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"strconv"
	"strings"
)

type ContractRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type Contract struct {
	Address string `json:"Address"`
	ABI string `json:"ABI"`
	Functions Functions `json:"HashableFunctions"`
	MaxMintable int `json:"MaxMintable"`
	MaxIncrement int `json:"MaxIncrement"`
	ContractOwner string `json:"ContractOwner"`
}

type dynamoContract struct {
	Address string `json:"Address"`
	ABI string `json:"ABI"`
	Functions string `json:"HashableFunctions"`
	MaxMintable int `json:"MaxMintable"`
	MaxIncrement int `json:"MaxIncrement"`
	ContractOwner string `json:"ContractOwner"`
}

func (c *Contract) fromDynamo(dContract *dynamoContract) error {
	functions := Functions{}
	err := json.Unmarshal([]byte(dContract.Functions), &functions)
	if err != nil {
		return err
	}
	c.Address = dContract.Address
	c.ABI = dContract.ABI
	c.Functions = functions
	c.MaxMintable = dContract.MaxMintable
	c.MaxIncrement = dContract.MaxIncrement
	c.ContractOwner = dContract.ContractOwner
	return nil
}

func (c *Contract) ToRPC() (*pb.Contract) {
	functions := &pb.Functions{Functions: map[string]*pb.Function{}}
	for key, val := range c.Functions.Functions {
		function := &pb.Function{}
		for _, arg := range val.Arguments {
			function.Arguments = append(function.Arguments, &pb.Argument{
				Name: arg.Name,
				Type: arg.Type,
			})
		}
		functions.Functions[key] = function
	}
	return &pb.Contract{
		Address:      c.Address,
		Abi:          c.ABI,
		HashableFunctions: 	  functions,
		MaxMintable:  int64(c.MaxMintable),
		MaxIncrement: int64(c.MaxIncrement),
		Owner:        c.ContractOwner,
	}
}

func (c *Contract) FromRPC(contract *pb.Contract) () {
	functions := Functions{Functions: map[string]Function{}}
	for key, val := range contract.HashableFunctions.Functions {
		function := Function{}
		for _, arg := range val.Arguments {
			function.Arguments = append(function.Arguments, Argument{
				Name: arg.Name,
				Type: arg.Type,
			})
		}
		functions.Functions[key] = function
	}
	c.Address = contract.Address
	c.ABI = contract.Abi
	c.Functions = functions
	c.MaxMintable = int(contract.MaxMintable)
	c.MaxIncrement = int(contract.MaxIncrement)
	c.ContractOwner = contract.Owner
}

func NewContractRepository(tableName string, cfg ...*aws.Config) (ContractRepository, error) {
	sess, err := session.NewSession(cfg...)
	if err != nil {
		return nil, err
	}
	repo := &ContractRepo{dynamodb.New(sess, cfg...), tableName}
	return repo, nil
}

func (cr *ContractRepo) Init() error {
	log.Println("Attempting to Create Table: " + cr.tableName)
	_, createErr := cr.db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions:   []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Address"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ContractOwner"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema:              []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Address"),
				KeyType: aws.String("HASH"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{
				IndexName: aws.String("ContractOwner"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("ContractOwner"),
						KeyType: aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					NonKeyAttributes: nil,
					ProjectionType:   aws.String("ALL"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits: aws.Int64(5),
					WriteCapacityUnits:  aws.Int64(5),
				},
			},
		},
		TableName:              aws.String(cr.tableName),
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(5),
			WriteCapacityUnits:  aws.Int64(5),
		},
	})

	log.Println("Table Creation request complete")

	if createErr == nil || strings.Contains(createErr.Error(), dynamodb.ErrCodeTableAlreadyExistsException) ||
		strings.Contains(createErr.Error(), dynamodb.ErrCodeGlobalTableAlreadyExistsException) ||
		strings.Contains(createErr.Error(), dynamodb.ErrCodeResourceInUseException) {
		return nil
	}

	return createErr
}

func (cr *ContractRepo) GetContract(ctx context.Context, contractAddress string) (*Contract, error) {
	result, err := cr.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Address": {
				S: aws.String(contractAddress),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.New("could not find key for contract: " + contractAddress)
	}
	contract := Contract{}
	dContract := dynamoContract{}


	if unmarshalErr := dynamodbattribute.UnmarshalMap(result.Item, &dContract); unmarshalErr != nil {
		return nil, unmarshalErr
	}
	if convertErr := contract.fromDynamo(&dContract); convertErr != nil {
		return nil, convertErr
	}
	return &contract, nil
}

func (cr *ContractRepo) GetContractsByOwner(ctx context.Context, owner string) ([]*Contract, error) {
	result, err := cr.db.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(cr.tableName),
		IndexName: aws.String("ContractOwner"),
		KeyConditionExpression: aws.String("ContractOwner = :v_owner"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v_owner": {S: aws.String(owner)},
		},
	})
	if err != nil {
		return nil, err
	}
	contracts := []*Contract{}
	for _, item := range result.Items {
		contract := Contract{}
		dContract := dynamoContract{}
		marshalErr := dynamodbattribute.UnmarshalMap(item, &dContract)
		if marshalErr != nil {
			log.Println(marshalErr.Error())
			continue
		}
		if convertErr := contract.fromDynamo(&dContract); convertErr != nil {
			log.Println(convertErr.Error())
			continue
		}
		contracts = append(contracts, &contract)
	}
	return contracts, nil
}


func (cr *ContractRepo) UpsertContract(ctx context.Context, contract *Contract) error {
	maxMint := strconv.FormatInt(int64(contract.MaxMintable), 10)
	maxIncr := strconv.FormatInt(int64(contract.MaxIncrement), 10)
	funcStr, marshalErr := json.Marshal(contract.Functions)
	if marshalErr != nil {
		return marshalErr
	}
	_, err := cr.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(cr.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Address": {
				S: aws.String(contract.Address),
			},
			"ABI": {
				S: aws.String(contract.ABI),
			},
			"HashableFunctions": {
				S: aws.String(string(funcStr)),
			},
			"MaxMintable": {
				N: aws.String(maxMint),
			},
			"MaxIncrement": {
				N: aws.String(maxIncr),
			},
			"ContractOwner": {
				S: aws.String(contract.ContractOwner),
			},
		},
	})
	return err
}
func (cr *ContractRepo) DeleteContract(ctx context.Context, contractAddress, owner string) error {
	_, err := cr.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Address": {
				S: aws.String(contractAddress),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":owner": {S: aws.String(owner)},
		},
		ConditionExpression: aws.String("ContractOwner = :owner"),
	})
	return err
}
