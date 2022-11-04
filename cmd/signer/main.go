package main

import (
	"contract-service/signing"
	"contract-service/utils"
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
	log.Printf("Loading SignatureHandler with Config: \n%+v\n", cfg)
	server, err := signing.InitializeSigningServer(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, cfg.TableName, &aws.Config{
		Endpoint:         aws.String(cfg.AWSEndpoint),
		Region:           aws.String(cfg.AWSRegion),
		Credentials:      credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		DisableSSL:       aws.Bool(!cfg.SSLEnabled),
	})
	if err != nil {
		log.Fatal(err)
	}
	liveProbeErr := make(chan string)
	probe := utils.NewProbe()

	probe.Serve(liveProbeErr)
	log.Fatal(utils.MergeChannels(liveProbeErr, server.Channel))
}

