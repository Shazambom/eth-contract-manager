package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"strings"
)

type PrivateKeyRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type ContractKeyPair struct {
	ContractAddress string `json:"ContractAddress"`
	PrivateKey string `json:"PrivateKey"`
}


func NewPrivateKeyRepository(tableName string, cfg ...*aws.Config) (PrivateKeyRepository, error) {
	sess, err := session.NewSession(cfg...)
	if err != nil {
		return nil, err
	}
	repo := &PrivateKeyRepo{dynamodb.New(sess, cfg...), tableName}
	return repo, nil
}

func (pkr *PrivateKeyRepo) Init() error {
	log.Println("Attempting to Create Table: " + pkr.tableName)
	_, createErr := pkr.db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions:   []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ContractAddress"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema:              []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("ContractAddress"), KeyType: aws.String("HASH")},
		},
		TableName:              aws.String(pkr.tableName),
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

// GetPrivateKey ---- WARNING ---- NEVER CALL OUTSIDE OF SIGNER SERVICE DUE TO SECURITY CONCERNS
func (pkr *PrivateKeyRepo) GetPrivateKey(ctx context.Context, contractAddress string) (string, error) {
	result, err := pkr.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(pkr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ContractAddress": {
				S: aws.String(contractAddress),
			},
		},
	})
	if err != nil {
		return "", err
	}
	if result.Item == nil {
		return "", errors.New("could not find key for contract: " + contractAddress)
	}
	keyPair := ContractKeyPair{}

	unmarshalErr := dynamodbattribute.UnmarshalMap(result.Item, &keyPair)
	if unmarshalErr != nil {
		return "", unmarshalErr
	}

	return keyPair.PrivateKey, nil
}

func (pkr *PrivateKeyRepo) UpsertPrivateKey(ctx context.Context, contractAddress, key string) error {
	_, err := pkr.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(pkr.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"ContractAddress": {
				S: aws.String(contractAddress),
			},
			"PrivateKey": {
				S: aws.String(key),
			},
		},
	})
	return err
}

func (pkr *PrivateKeyRepo) DeletePrivateKey(ctx context.Context, contractAddress string) error {
	_, err := pkr.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(pkr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ContractAddress": {
				S: aws.String(contractAddress),
			},
		},
	})
	return err
}