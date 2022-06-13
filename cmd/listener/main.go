package main

import (
	"contract-service/listener"
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
)
//TODO Add Ping route to container to check if service is alive
func main() {
	log.Println("Getting environment variables")
	cfg, envErr := NewConfig()
	if envErr != nil {
		log.Fatal(envErr)
	}
	log.Printf("Loading Listener with Config: \n%+v\n", cfg)
	handler, handlerInitErr := listener.InitializeListenerService(&aws.Config{
		Endpoint: aws.String(cfg.AWSEndpoint),
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL: aws.Bool(!cfg.SSLEnabled),
	}, cfg.TableName)
	if handlerInitErr != nil {
		log.Fatal(handlerInitErr)
	}
	log.Println("Prepping handler dependencies")
	if handlerDepErr := handler.InitService(); handlerDepErr != nil {
		log.Fatal(handlerDepErr)
	}
	rds := storage.NewRedisListener(storage.RedisConfig{Endpoint: cfg.RedisEndpoint, Password: cfg.RedisPassword})
	defer rds.Close()
	log.Println("Initializing events")
	if initErr := rds.InitEvents(); initErr != nil {
		log.Fatal(initErr)
	}
	log.Println("Starting to listen to Redis event stream")
	if err := rds.Listen(handler.Handle); err != nil {
		log.Fatal(err)
	}
}
