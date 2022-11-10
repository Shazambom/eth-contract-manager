package main

import (
	"contract-service/contracts"
	"contract-service/storage"
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
	log.Printf("Loading ContractManager with Config: \n%s\n", cfg.String())
	contractRPC, contractErr := contracts.InitializeContractServer(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, storage.ContractConfig{ TableName: cfg.TableName,  CFG: []*aws.Config{{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(!cfg.SSLEnabled),
	}}})
	if contractErr != nil {
		log.Fatal(contractErr)
	}
	log.Fatal(<-contractRPC.Channel)
}
