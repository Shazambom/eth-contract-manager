package main

import (
	"contract-service/signing"
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
	log.Printf("Loading Signer with Config: \n%+v\n", cfg)
	server, err := signing.InitializeSigningServer(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, cfg.TableName, &aws.Config{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(!cfg.SSLEnabled),
	})
	if err != nil {
		log.Fatal(err)
	}
	errorCode := <-server.Channel
	log.Println(errorCode)
}

