package storage

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
)

type TransactionRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

type TransactionConfig struct {
	tableName string
	cfg []*aws.Config
}

func NewTransactionRepo(config TransactionConfig) (TransactionRepository, error) {
	sess, err := session.NewSession(config.cfg...)
	if err != nil {
		return nil, err
	}
	repo := &TransactionRepo{dynamodb.New(sess, config.cfg...), config.tableName}
	return repo, nil
}

func (tr *TransactionRepo) StoreTransaction(ctx context.Context, token Token) error {
	_, err := tr.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"abi": {S: aws.String(token.ABI)},
			"abi_packed_txn": {B: token.ABIPackedTxn},
			"contract_address": {S: aws.String(token.ContractAddress)},
			"user_address": {S: aws.String(token.UserAddress)},
			"hash": {S: aws.String(token.Hash)},
		},
		TableName: aws.String(tr.tableName),
	})
	return err
}

func (tr *TransactionRepo) GetTransactions(ctx context.Context, address string) ([]*Token, error) {
	result, err := tr.db.QueryWithContext(ctx, &dynamodb.QueryInput{
		TableName: aws.String(tr.tableName),
		IndexName: aws.String("user_address"),
		KeyConditionExpression: aws.String("user_address = :v_addr"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v_addr": {S: aws.String(address)},
		},
	})
	if err != nil {
		return nil, err
	}
	tokens := []*Token{}
	for _, item := range result.Items {
		token := Token{}
		marshalErr := dynamodbattribute.UnmarshalMap(item, &token)
		if marshalErr != nil {
			log.Println(marshalErr.Error())
			continue
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}

func (tr *TransactionRepo) DeleteTransaction(ctx context.Context, address, hash string) error {
	_, err := tr.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(tr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_address": {S: aws.String(address)},
			"hash": {S: aws.String(hash)},
		},
	})
	return err
}
