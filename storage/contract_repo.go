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
)

type ContractRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type ContractConfig struct {
	TableName string
	CFG []*aws.Config
}

type Contract struct {
	Address string `json:"Address"`
	ABI string `json:"ABI"`
	Functions map[string]Function `json:"Functions"`
	ContractOwner string `json:"ContractOwner"`
}

type dynamoContract struct {
	Address string `json:"Address"`
	ABI string `json:"ABI"`
	Functions string `json:"Functions"`
	ContractOwner string `json:"ContractOwner"`
}

func (c *Contract) fromDynamo(dContract *dynamoContract) error {
	functions := map[string]Function{}
	err := json.Unmarshal([]byte(dContract.Functions), &functions)
	if err != nil {
		return err
	}
	c.Address = dContract.Address
	c.ABI = dContract.ABI
	c.Functions = functions
	c.ContractOwner = dContract.ContractOwner
	return nil
}

func (c *Contract) ToRPC() (*pb.Contract) {
	functions := map[string]*pb.Function{}
	for key, val := range c.Functions {
		function := &pb.Function{Arguments: []*pb.Argument{}}
		for _, arg := range val.Arguments {
			function.Arguments = append(function.Arguments, &pb.Argument{
				Name: arg.Name,
				Type: arg.Type,
			})
		}
		functions[key] = function
	}
	return &pb.Contract{
		Address:      c.Address,
		Abi:          c.ABI,
		Functions: 	  functions,
		Owner:        c.ContractOwner,
	}
}

func (c *Contract) FromRPC(contract *pb.Contract) () {
	functions := map[string]Function{}
	if contract.Functions != nil {
		for key, val := range contract.Functions{
			function := Function{Arguments: []Argument{}}
			for _, arg := range val.Arguments {
				function.Arguments = append(function.Arguments, Argument{
					Name: arg.Name,
					Type: arg.Type,
				})
			}
			functions[key] = function
		}
	}
	c.Address = contract.Address
	c.ABI = contract.Abi
	c.Functions = functions
	c.ContractOwner = contract.Owner
}

func NewContractRepository(config ContractConfig) (ContractRepository, error) {
	sess, err := session.NewSession(config.CFG...)
	if err != nil {
		return nil, err
	}
	repo := &ContractRepo{dynamodb.New(sess, config.CFG...), config.TableName}
	return repo, nil
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
			"Functions": {
				S: aws.String(string(funcStr)),
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
