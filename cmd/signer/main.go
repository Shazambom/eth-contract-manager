package main

import (
	"bitbucket.org/artie_inc/contract-service/signing"
	"bitbucket.org/artie_inc/contract-service/storage"
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
	log.Printf("Loading SignatureHandler with Config: \n%s\n", cfg.String())
	awsCredentials := credentials.NewEnvCredentials()
	if cfg.AccessKeyID != "" && cfg.SecretAccessKey != "" {
		awsCredentials = credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	}
	server, err := signing.InitializeSigningServer(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, storage.PrivateKeyConfig{
		TableName: cfg.TableName,
		CFG: []*aws.Config{{
			Endpoint:    aws.String(cfg.AWSEndpoint),
			Region:      aws.String(cfg.AWSRegion),
			Credentials: awsCredentials,
			DisableSSL:  aws.Bool(!cfg.SSLEnabled),
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(<-server.Channel)
}
