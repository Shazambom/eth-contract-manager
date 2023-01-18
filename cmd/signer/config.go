package main

import (
	"contract-service/utils"
	"fmt"
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
}

func NewConfig() (Config, error) {
	var port, awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, tableName string
	envErr := utils.GetEnvVarBatch([]string{"PORT", "AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "TABLE_NAME"}, &port, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &tableName)
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
	}, envErr
}

func (c *Config) String() string {
	return fmt.Sprintf("{\n\tPort: %d\n\tTableName: %s\n\tAWSEndpoint: %s\n\tAWSRegion: %s\n\tAccessKeyID: ********%s\n\tSecretAccessKey: ********%s\n\tSSLEnabled: %t\n}\n",
		c.Port, c.TableName, c.AWSEndpoint, c.AWSRegion, c.AccessKeyID[len(c.AccessKeyID)-3:], c.SecretAccessKey[len(c.SecretAccessKey)-3:], c.SSLEnabled)
}
