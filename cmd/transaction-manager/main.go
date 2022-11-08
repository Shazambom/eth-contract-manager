package main

import (
	"contract-service/contracts"
	"contract-service/signing"
	"contract-service/storage"
	"contract-service/utils"
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
	log.Printf("Loading TransactionManager with Config: \n%+v\n", cfg)
	signingClient, clientErr := signing.NewClient(cfg.SignerEndpoint, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	awsConfig := &aws.Config{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(!cfg.SSLEnabled),
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
	liveProbeErr := make(chan string)
	probe := utils.NewProbe()

	probe.Serve(liveProbeErr)
	log.Fatal(utils.MergeChannels(liveProbeErr, transactionRPC.Channel))
}
