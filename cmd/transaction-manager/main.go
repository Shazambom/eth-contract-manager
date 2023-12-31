package main

import (
	"bitbucket.org/artie_inc/contract-service/contracts"
	"bitbucket.org/artie_inc/contract-service/signing"
	"bitbucket.org/artie_inc/contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	log.Printf("Loading TransactionManager with Config: \n%s\n", cfg.String())
	signingClient, clientErr := signing.NewClient(cfg.SignerEndpoint, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	awsCredentials := credentials.NewEnvCredentials()
	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		awsCredentials = credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	}
	awsConfig := &aws.Config{
		Endpoint:    aws.String(cfg.AWSEndpoint),
		Region:      aws.String(cfg.AWSRegion),
		Credentials: awsCredentials,
		DisableSSL:  aws.Bool(!cfg.SSLEnabled),
	}

	transactionRPC, gRPCErr := contracts.InitializeTransactionServer(
		cfg.Port,
		[]grpc.ServerOption{grpc.EmptyServerOption{}},
		signingClient.SigningClient,
		storage.ContractConfig{
			TableName: cfg.ContractTableName,
			CFG:       []*aws.Config{awsConfig},
		},
		storage.TransactionConfig{
			TableName: cfg.TransactionTableName,
			CFG:       []*aws.Config{awsConfig},
		})
	if gRPCErr != nil {
		log.Fatal(gRPCErr)
	}
	log.Fatal(<-transactionRPC.Channel)
}
