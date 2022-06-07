package main

import (
	"contract-service/contracts"
	"contract-service/storage"
	"log"
)

//TODO Implement config management
//TODO Implement wire dependency injection

func main() {
	contractRepo := storage.NewContractRepository("Contracts", nil, nil)

	contractHandler := contracts.NewContractManagerHandler(contractRepo)

	contractgRPC, contractErr := contracts.NewContractServer(8082, nil, contractHandler)
	if contractErr != nil {
		log.Fatal(contractErr)
	}

	errorCode := <-contractgRPC.Channel
	log.Println(errorCode)
}
