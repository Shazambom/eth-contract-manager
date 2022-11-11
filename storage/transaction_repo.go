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
	TableName string
	CFG []*aws.Config
}

func NewTransactionRepo(config TransactionConfig) (TransactionRepository, error) {
	sess, err := session.NewSession(config.CFG...)
	if err != nil {
		return nil, err
	}
	repo := &TransactionRepo{dynamodb.New(sess, config.CFG...), config.TableName}
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
			"is_complete": {BOOL: aws.Bool(token.IsComplete)},
		},
		TableName: aws.String(tr.tableName),
	})
	return err
}

func (tr *TransactionRepo) queryTransactionsTable(ctx context.Context, input *dynamodb.QueryInput) ([]*Token, error) {
	result, err := tr.db.QueryWithContext(ctx, input)
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

func (tr *TransactionRepo) GetTransactions(ctx context.Context, address string) ([]*Token, error) {
	return tr.queryTransactionsTable(ctx, &dynamodb.QueryInput{
		TableName: aws.String(tr.tableName),
		IndexName: aws.String("user_address"),
		KeyConditionExpression: aws.String("user_address = :v_addr"),
		FilterExpression: aws.String("is_complete = :ic"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v_addr": {S: aws.String(address)},
			":ic": {BOOL: aws.Bool(false)},
		},
	})
}


func (tr *TransactionRepo) GetAllTransactions(ctx context.Context, address string) ([]*Token, error) {
	return tr.queryTransactionsTable(ctx, &dynamodb.QueryInput{
		TableName: aws.String(tr.tableName),
		IndexName: aws.String("user_address"),
		KeyConditionExpression: aws.String("user_address = :v_addr"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":v_addr": {S: aws.String(address)},
		},
	})
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

func (tr *TransactionRepo) CompleteTransaction(ctx context.Context, address, hash string) error {
	_, err := tr.db.UpdateItemWithContext(ctx, &dynamodb.UpdateItemInput{
		UpdateExpression:            aws.String("SET is_complete = :ic"),
		ExpressionAttributeValues:   map[string]*dynamodb.AttributeValue{
			":ic": {BOOL: aws.Bool(true)},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"user_address": {S: aws.String(address)},
			"hash": {S: aws.String(hash)},
		},
		TableName:                   aws.String(tr.tableName),
	})
	return err
}