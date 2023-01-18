package main

import (
	"contract-service/api"
	"contract-service/contracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//TODO Move config stuff to the config struct and implement dependency injection with wire
	txnClient, clientErr := contracts.NewTransactionClient("transaction-manager:8083", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}
	claimService, claimErr := api.NewContractIntegrationRPCService(8085, []grpc.ServerOption{grpc.EmptyServerOption{}}, txnClient)
	if claimErr != nil {
		log.Fatal(claimErr)
	}
	log.Fatal(<-claimService.Channel)
}

