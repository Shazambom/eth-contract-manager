package main

//TODO Implement config management
//TODO Implement wire dependency injection

import (
	"contract-service/listener"
	"contract-service/storage"
	"contract-service/utils"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
)
//TODO Clean up env variable management & struct initialization. Move struct initialization to wire
func main() {
	fmt.Println("Getting environment variables")
	var awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, bucketName string
	envErr := utils.GetEnvVarBatch([]string{"AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "BUCKET_NAME"}, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &bucketName)
	if envErr != nil {
		log.Fatal(envErr)
	}
	fmt.Println("Building S3 connection")
	var s3, s3Err = storage.NewS3(&aws.Config{
		Endpoint: aws.String(awsEndpoint),
		Region: aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsKeyId, awsSecret, ""),
		S3ForcePathStyle: aws.Bool(true),
		DisableSSL: aws.Bool(sslEnabled != "true"),
	}, bucketName)
	if s3Err != nil {
		log.Fatal(s3Err)
	}
	if s3initErr := s3.InitBucket(); s3initErr != nil {
		log.Fatal(s3initErr)
	}
	fmt.Println("Connecting to Redis client")
	rds := storage.NewRedisListener(os.Getenv("RDS_ENDPOINT"), os.Getenv("RDS_PWD"))
	defer rds.Close()
	fmt.Println("Initializing events")
	if initErr := rds.InitEvents(); initErr != nil {
		fmt.Println(initErr)
	}
	fmt.Println("Building EventHandlerService")
	handler := listener.NewEventHandlerService(s3)
	fmt.Println("Starting to listen to Redis event stream")
	if err := rds.Listen(handler.Handle); err != nil {
		log.Fatal(err)
	}
}
