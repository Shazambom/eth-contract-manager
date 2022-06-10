package storage

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func TableExists(db *dynamodb.DynamoDB, tableName string) error {
	out, lsErr := db.ListTables(&dynamodb.ListTablesInput{})
	if lsErr != nil {
		return lsErr
	}
	for _, table := range out.TableNames {
		log.Println(table)
		if *table == tableName {
			return nil
		}
	}
	return errors.New("couldn't find table: " + tableName)
}
