package main

import (
	"contract-service/utils"
	"fmt"
	"strconv"
)

type Config struct {
	TxnHost string
	Port int
	AlivePort int
}

func NewConfig() (Config, error) {
	var port, alive, txnHost string
	envErr := utils.GetEnvVarBatch([]string{"PORT", "ALIVE_PORT", "TXN_HOST"}, &port, &alive, &txnHost)
	if envErr != nil {
		return Config{}, envErr
	}
	prt, convErr := strconv.Atoi(port)
	if convErr != nil {
		return Config{}, convErr
	}
	alivePort, aliveErr := strconv.Atoi(alive)
	if aliveErr != nil {
		return Config{}, aliveErr
	}
	return Config{
		TxnHost: txnHost,
		Port:    prt,
		AlivePort: alivePort,
	}, nil
}

func (c *Config) String()string {
	return fmt.Sprintf("{\n\tPort: %d\n\tAlivePort: %d\n\tTxnHost: %s}\n",
		c.Port, c.AlivePort, c.TxnHost)
}