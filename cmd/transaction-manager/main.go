package main

import (
	"contract-service/contracts"
	"contract-service/signing"
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"google.golang.org/grpc"
	"log"
)

//TODO Add Ping route to container to check if service is alive

func main() {
	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	signingClient, clientErr := signing.NewClient(cfg.SignerEndpoint, []grpc.DialOption{grpc.EmptyDialOption{}})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	transactionRPC, gRPCErr := contracts.InitializeTransactionServer(
		cfg.Port,
		nil,
		storage.RedisConfig{
			Endpoint: cfg.RedisEndpoint,
			Password: cfg.RedisPwd,
			CountKey: cfg.CountKey,
		},
		signingClient.SigningClient,
		cfg.TableName,
		&aws.Config{
			Endpoint:         aws.String(cfg.AWSEndpoint),
			Region:           aws.String(cfg.AWSRegion),
			Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
			DisableSSL:       aws.Bool(cfg.SSLEnabled),
	})
	if gRPCErr != nil {
		log.Fatal(gRPCErr)
	}
	errorCode := <- transactionRPC.Channel
	log.Println(errorCode)
}
