package main

import (
	"contract-service/contracts"
	"contract-service/web"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	errChan := make(chan string)
	probe := web.NewProbe()


	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	txnClient, clientErr := contracts.NewTransactionClient(cfg.TxnHost, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	transactionAPI := web.NewTransactionAPI(txnClient)

	transactionAPI.Serve(cfg.Port, errChan)
	probe.Serve(cfg.AlivePort, errChan)
	log.Fatal(<-errChan)
}

