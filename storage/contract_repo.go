package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"strconv"
)

type ContractRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type Contract struct {
	Address string `json:"Address"`
	ABI string `json:"ABI"`
	Functions []*string `json:"Functions"`
	MaxMintable int `json:"MaxMintable"`
	MaxIncrement int `json:"MaxIncrement"`
	Owner string `json:"Owner"`
}

func NewContractRepository(tableName string, sess *session.Session, cfg ...*aws.Config) ContractRepository {
	return &ContractRepo{dynamodb.New(sess, cfg...), tableName}
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

	unmarshalErr := dynamodbattribute.UnmarshalMap(result.Item, &contract)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return &contract, nil
}

func (cr *ContractRepo) GetContractsByOwner(ctx context.Context, owner string) ([]*Contract, error) {
	result, err := cr.db.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(cr.tableName),
		IndexName: aws.String("Owner"),
		KeyConditionExpression: aws.String("Owner = :v_owner"),
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
		marshalErr := dynamodbattribute.UnmarshalMap(item, &contract)
		if marshalErr != nil {
			log.Println(marshalErr.Error())
			continue
		}
		contracts = append(contracts, &contract)
	}
	return contracts, nil
}


func (cr *ContractRepo) UpsertContract(ctx context.Context, contract *Contract) error {
	maxMint := strconv.FormatInt(int64(contract.MaxMintable), 10)
	maxIncr := strconv.FormatInt(int64(contract.MaxIncrement), 10)
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
				SS: contract.Functions,
			},
			"MaxMintable": {
				N: aws.String(maxMint),
			},
			"MaxIncrement": {
				N: aws.String(maxIncr),
			},
			"Owner": {
				S: aws.String(contract.Owner),
			},
		},
	})
	return err
}
func (cr *ContractRepo) DeleteContract(ctx context.Context, contractAddress string) error {
	_, err := cr.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(cr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Address": {
				S: aws.String(contractAddress),
			},
		},
	})
	return err
}
