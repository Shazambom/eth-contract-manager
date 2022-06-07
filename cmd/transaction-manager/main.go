package main

import (
	"contract-service/contracts"
	"contract-service/signing"
	"contract-service/storage"
	"log"
)

//TODO Implement config management
//TODO Implement wire dependency injection

func main() {
	contractRepo := storage.NewContractRepository("Contracts", nil, nil)

	signingClient, clientErr := signing.NewClient("signer:8081", []grpc.DialOption{grpc.WithInsecure()})
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	writer := storage.NewRedisWriter("redis:9999", "", "Count")

	transactionHandler := contracts.NewContractTransactionHandler(writer, contractRepo, signingClient.SigningClient)

	transactiongRPC, gRPCErr := contracts.NewTransactionServer(8083, nil, transactionHandler)
	if gRPCErr != nil {
		log.Fatal(gRPCErr)
	}
	errorCode := <- transactiongRPC.Channel
	log.Println(errorCode)
}
