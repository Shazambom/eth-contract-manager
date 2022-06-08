package main


import (
	"contract-service/listener"
	"contract-service/storage"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
)
//TODO Add Ping route to container to check if service is alive
func main() {
	fmt.Println("Getting environment variables")
	cfg, envErr := NewConfig()
	if envErr != nil {
		log.Fatal(envErr)
	}
	rds := storage.NewRedisListener(os.Getenv("RDS_ENDPOINT"), os.Getenv("RDS_PWD"))
	defer rds.Close()
	fmt.Println("Initializing events")
	if initErr := rds.InitEvents(); initErr != nil {
		fmt.Println(initErr)
	}
	handler, handlerInitErr := listener.InitializeListenerService(&aws.Config{
		Endpoint: aws.String(cfg.AWSEndpoint),
		Region: aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL: aws.Bool(cfg.SSLEnabled),
	}, cfg.TableName)
	if handlerInitErr != nil {
		log.Fatal(handlerInitErr)
	}
	fmt.Println("Starting to listen to Redis event stream")
	if err := rds.Listen(handler.Handle); err != nil {
		log.Fatal(err)
	}
}
