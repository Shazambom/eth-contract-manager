package main

import "contract-service/utils"

type Config struct {
	TableName string
	AWSEndpoint string
	AWSRegion string
	AccessKeyID string
	SecretAccessKey string
	SSLEnabled bool
}

func NewConfig() (Config, error) {
	var awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, bucketName string
	envErr := utils.GetEnvVarBatch([]string{"AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "BUCKET_NAME"}, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &bucketName)
	if envErr != nil {
		return Config{}, envErr
	}
	return Config{
		TableName: bucketName,
		AWSEndpoint: awsEndpoint,
		AWSRegion: awsRegion,
		AccessKeyID: awsKeyId,
		SecretAccessKey: awsSecret,
		SSLEnabled: sslEnabled == "true",
	}, envErr
}
