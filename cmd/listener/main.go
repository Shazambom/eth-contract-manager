package main

import (
	"contract-service/storage"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"
	"os"
)

var s3, s3Err = storage.NewS3(&aws.Config{
	Endpoint: aws.String(os.Getenv("AWS_ENDPOINT")),
	Region: aws.String(os.Getenv("AWS_REGION")),
	Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	S3ForcePathStyle: aws.Bool(true),
	DisableSSL: aws.Bool(os.Getenv("SSL_ENABLED") != "true"),
},
os.Getenv("BUCKET_NAME"))

func EventHandler(key, val string, err error) error {
	if err != nil {
		fmt.Printf("Error with redis stream: %s\n", err)
		return err
	}
	fmt.Printf("key: %s\nval:%s\n", key, val)
	token := storage.Token{}
	parseErr := json.Unmarshal([]byte(val), &token)
	if parseErr != nil {
		fmt.Printf("Error with redis format: %s\n", parseErr.Error())
		return nil
	}
	storeErr := s3.StoreToken(&token)
	if storeErr != nil {
		fmt.Printf("Error storing in s3: %s\n", storeErr.Error())
	}
	return nil
}


func main() {
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
	fmt.Println("Starting to listen to Redis event stream")
	if err := rds.Listen(EventHandler); err != nil {
		log.Fatal(err)
	}
}
