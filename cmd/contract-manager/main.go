package main

import (
	"contract-service/contracts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	contractRPC, contractErr := contracts.InitializeContractServer(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, cfg.TableName, &aws.Config{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(cfg.SSLEnabled),
	})
	if contractErr != nil {
		log.Fatal(contractErr)
	}

	errorCode := <-contractRPC.Channel
	log.Println(errorCode)
}
