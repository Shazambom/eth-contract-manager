package main

import (
"contract-service/utils"
"fmt"
"strconv"
)

type Config struct {
	TxnHost string
	Port int
}

func NewConfig() (Config, error) {
	var port, txnHost string
	envErr := utils.GetEnvVarBatch([]string{"PORT", "TXN_HOST"}, &port, &txnHost)
	if envErr != nil {
		return Config{}, envErr
	}
	prt, convErr := strconv.Atoi(port)
	if convErr != nil {
		return Config{}, convErr
	}
	return Config{
		TxnHost: txnHost,
		Port:    prt,
	}, nil
}

func (c *Config) String()string {
	return fmt.Sprintf("{\n\tPort: %d\n\tTxnHost: %s}\n",
		c.Port, c.TxnHost)
}
