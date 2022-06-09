package main

import (
	"contract-service/utils"
	"strconv"
)

type Config struct {
	Port int
	TableName string
	AWSEndpoint string
	AWSRegion string
	AccessKeyID string
	SecretAccessKey string
	SSLEnabled bool
	RedisEndpoint string
	RedisPwd string
	CountKey string
	SignerEndpoint string
}

func NewConfig() (Config, error) {
	var port, awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, tableName, redisEndpoint, redisPwd, countKey, signerEndpoint string
	envErr := utils.GetEnvVarBatch([]string{"PORT", "AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "TABLE_NAME", "REDIS_ENDPOINT", "REDIS_PASSWORD", "COUNT_KEY", "SIGNER_HOST"}, &port, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &tableName, &redisEndpoint, &redisPwd, &countKey, &signerEndpoint)
	if envErr != nil {
		return Config{}, envErr
	}
	prt, convErr := strconv.Atoi(port)
	if convErr != nil {
		return Config{}, convErr
	}
	return Config{
		Port: prt,
		TableName: tableName,
		AWSEndpoint: awsEndpoint,
		AWSRegion: awsRegion,
		AccessKeyID: awsKeyId,
		SecretAccessKey: awsSecret,
		SSLEnabled: sslEnabled == "true",
		RedisEndpoint: redisEndpoint,
		RedisPwd: redisPwd,
		CountKey: countKey,
		SignerEndpoint: signerEndpoint,
	}, envErr
}
