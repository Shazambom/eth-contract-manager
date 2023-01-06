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

	//TODO Move config stuff to the config struct and implement dependency injection with wire

	txnClient, clientErr := contracts.NewTransactionClient("transaction-manager:8083", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	transactionAPI := web.NewTransactionAPI(txnClient)

	transactionAPI.Serve(8084, errChan)
	probe.Serve(8080, errChan)
	log.Fatal(<-errChan)
}

