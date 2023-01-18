package storage

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type PrivateKeyRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type PrivateKeyConfig struct {
	TableName string
	CFG []*aws.Config
}

type ContractKeyPair struct {
	ContractAddress string `json:"ContractAddress"`
	PrivateKey string `json:"PrivateKey"`
}


func NewPrivateKeyRepository(config PrivateKeyConfig) (PrivateKeyRepository, error) {
	sess, err := session.NewSession(config.CFG...)
	if err != nil {
		return nil, err
	}
	repo := &PrivateKeyRepo{dynamodb.New(sess, config.CFG...), config.TableName}
	return repo, nil
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