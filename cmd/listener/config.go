package main

import "contract-service/utils"

type Config struct {
	TableName string
	AWSEndpoint string
	AWSRegion string
	AccessKeyID string
	SecretAccessKey string
	SSLEnabled bool
	RedisEndpoint string
	RedisPassword string
}

func NewConfig() (Config, error) {
	var awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, bucketName, redisEndpoint, redisPassword string
	envErr := utils.GetEnvVarBatch([]string{"AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "BUCKET_NAME", "REDIS_ENDPOINT", "REDIS_PASSWORD"}, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &bucketName, &redisEndpoint, &redisPassword)
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
		RedisEndpoint: redisEndpoint,
		RedisPassword: redisPassword,
	}, envErr
}
