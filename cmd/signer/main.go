package main

import (
	"contract-service/signing"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
)

//TODO Add Ping route to container to check if service is alive


func main() {
	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	server, err := signing.InitializeSigningServer(cfg.Port, nil, cfg.TableName, &aws.Config{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(cfg.SSLEnabled),
	})
	if err != nil {
		log.Fatal(err)
	}
	errorCode := <-server.Channel
	log.Println(errorCode)
}

