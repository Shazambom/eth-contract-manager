package main

import (
	"bitbucket.org/artie_inc/contract-service/utils"
	"fmt"
	"strconv"
)

type Config struct {
	Port                 int
	ContractTableName    string
	TransactionTableName string
	AWSEndpoint          string
	AWSRegion            string
	AccessKeyID          string
	SecretAccessKey      string
	SSLEnabled           bool
	SignerEndpoint       string
}

func NewConfig() (Config, error) {
	var port, awsEndpoint, awsRegion, awsKeyId, awsSecret, sslEnabled, contractTableName, transactionTableName, signerEndpoint string
	envErr := utils.GetEnvVarBatch([]string{"PORT", "AWS_ENDPOINT", "AWS_REGION", "AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "SSL_ENABLED", "CONTRACT_TABLE_NAME", "TRANSACTION_TABLE_NAME", "SIGNER_HOST"}, &port, &awsEndpoint, &awsRegion, &awsKeyId, &awsSecret, &sslEnabled, &contractTableName, &transactionTableName, &signerEndpoint)
	if envErr != nil {
		return Config{}, envErr
	}
	prt, convErr := strconv.Atoi(port)
	if convErr != nil {
		return Config{}, convErr
	}
	return Config{
		Port:                 prt,
		ContractTableName:    contractTableName,
		TransactionTableName: transactionTableName,
		AWSEndpoint:          awsEndpoint,
		AWSRegion:            awsRegion,
		AccessKeyID:          awsKeyId,
		SecretAccessKey:      awsSecret,
		SSLEnabled:           sslEnabled == "true",
		SignerEndpoint:       signerEndpoint,
	}, envErr
}

func (c *Config) String() string {
	return fmt.Sprintf("{\n\tPort: %d\n\tContractTableName: %s\n\tTransactionTableName: %s\n\tAWSEndpoint: %s\n\tAWSRegion: %s\n\tAccessKeyID: ********%s\n\tSecretAccessKey: ********%s\n\tSSLEnabled: %t\n\tSignerEndpoint: %s\n}\n",
		c.Port, c.ContractTableName, c.TransactionTableName, c.AWSEndpoint, c.AWSRegion, c.AccessKeyID[len(c.AccessKeyID)-3:], c.SecretAccessKey[len(c.SecretAccessKey)-3:], c.SSLEnabled, c.SignerEndpoint)
}
