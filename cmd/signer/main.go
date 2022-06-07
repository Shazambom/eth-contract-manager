package main

import (
	"contract-service/signing"
	"contract-service/storage"
	"log"
)

//TODO Implement config management
//TODO Implement wire dependency injection


func main() {
	repo := storage.NewPrivateKeyRepository("ContractPrivateKeyRepository", nil, nil)
	server, err := signing.NewSignerServer(8081, nil, signing.NewSigningService(), repo)
	if err != nil {
		log.Fatal(err)
	}
	errorCode := <-server.Channel
	log.Println(errorCode)
}

