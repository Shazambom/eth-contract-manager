package storage

import (
	"bitbucket.org/artie_inc/contract-service/utils"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
	"testing"
)

var dynamoCfg = &aws.Config{
	Endpoint:    aws.String(utils.GetEnvVarWithDefault("TEST_DYANMO_ENDPOINT", "localhost:8000")),
	Region:      aws.String(utils.GetEnvVarWithDefault("TEST_AWS_REGION", "us-east-1")),
	Credentials: credentials.NewStaticCredentials(utils.GetEnvVarWithDefault("TEST_AWS_ACCESS_KEY_ID", "xxx"), utils.GetEnvVarWithDefault("TEST_AWS_SECRET_ACCESS_KEY", "yyy"), ""),
	DisableSSL:  aws.Bool(true),
}

var ctx = context.Background()

var s3cfg = &aws.Config{
	Endpoint:         aws.String(utils.GetEnvVarWithDefault("TEST_S3_ENDPOINT", "localhost:4566")),
	Region:           aws.String(utils.GetEnvVarWithDefault("TEST_AWS_REGION", "us-east-1")),
	Credentials:      credentials.NewStaticCredentials(utils.GetEnvVarWithDefault("TEST_AWS_ACCESS_KEY_ID", "xxx"), utils.GetEnvVarWithDefault("TEST_AWS_SECRET_ACCESS_KEY", "yyy"), ""),
	S3ForcePathStyle: aws.Bool(true),
	DisableSSL:       aws.Bool(true),
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
