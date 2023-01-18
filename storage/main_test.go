package storage

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
	"testing"
)

var dynamoCfg = &aws.Config{
	Endpoint:         aws.String("localhost:8000"),
	Region:           aws.String("us-east-1"),
	Credentials:      credentials.NewStaticCredentials("xxx","yyy", ""),
	DisableSSL:       aws.Bool(true),
}

var rdsCfg = RedisConfig{
	Endpoint: "localhost:6379",
	Password: "pass",
	CountKey: "counterTest",
}
var ctx = context.Background()

var s3cfg = &aws.Config{
	Endpoint: aws.String("localhost:4566"),
	Region: aws.String("us-east-1"),
	Credentials: credentials.NewStaticCredentials("xxx", "yyy", ""),
	S3ForcePathStyle: aws.Bool(true),
	DisableSSL: aws.Bool(true),
}
var testBucketName = "buckety"

var PrivateKeyRepoConfig = PrivateKeyConfig{
	TableName: "ContractPrivateKeyRepository",
	CFG:       []*aws.Config{dynamoCfg},
}

var ContractRepoConfig = ContractConfig{
	TableName: "Contracts",
	CFG:       []*aws.Config{dynamoCfg},
}

var TransactionRepoConfig = TransactionConfig{
	TableName: "Transactions",
	CFG:       []*aws.Config{dynamoCfg},
}

var cr ContractRepository
var pkr PrivateKeyRepository
var tr TransactionRepository

func TestMain(m *testing.M) {
	s3, s3Err := NewS3(s3cfg, testBucketName)
	if s3Err != nil {
		log.Fatal(s3Err)
	}
	if s3InitErr := s3.InitBucket(); s3InitErr != nil {
		log.Fatal(s3InitErr)
	}

	var pkrErr error
	pkr, pkrErr = NewPrivateKeyRepository(PrivateKeyRepoConfig)
	if pkrErr != nil {
		log.Fatal(pkrErr)
	}
	var crErr error
	cr, crErr = NewContractRepository(ContractRepoConfig)
	if crErr != nil {
		log.Fatal(crErr)
	}
	var trErr error
	tr, trErr = NewTransactionRepo(TransactionRepoConfig)
	if trErr != nil {
		log.Fatal(trErr)
	}

	code := m.Run()

	os.Exit(code)
}
