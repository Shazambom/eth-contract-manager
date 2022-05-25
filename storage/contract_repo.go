package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type ContractRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type Contract struct {
	Address string `json:"Address"`
	ABI []*string `json:"ABI"`
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
func (cr *ContractRepo) UpsertContract(ctx context.Context, contract *Contract) error {
	_, err := cr.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(cr.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"Address": {
				S: aws.String(contract.Address),
			},
			"ABI": {
				SS: contract.ABI,
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
