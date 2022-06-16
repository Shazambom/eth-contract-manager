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

var TestContractTableName = "ContractsTest"
var TestPrivateKeyTableName = "PrivateKeyTest"

func TestMain(m *testing.M) {
	s3, s3Err := NewS3(s3cfg, testBucketName)
	if s3Err != nil {
		log.Fatal(s3Err)
	}
	if s3InitErr := s3.InitBucket(); s3InitErr != nil {
		log.Fatal(s3InitErr)
	}

	pkr, pkrErr := NewPrivateKeyRepository(TestPrivateKeyTableName, dynamoCfg)
	if pkrErr != nil {
		log.Fatal(pkrErr)
	}
	if pkrInitErr := pkr.Init(); pkrInitErr != nil {
		log.Fatal(pkrInitErr)
	}

	cr, crErr := NewContractRepository(TestContractTableName, dynamoCfg)
	if crErr != nil {
		log.Fatal(crErr)
	}
	if crInitErr := cr.Init(); crInitErr != nil {
		log.Fatal(crInitErr)
	}

	code := m.Run()

	os.Exit(code)
}
