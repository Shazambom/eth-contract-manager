package main

import (
	"bitbucket.org/artie_inc/contract-service/api"
	"bitbucket.org/artie_inc/contract-service/contracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cfg, cfgErr := NewConfig()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}
	txnClient, clientErr := contracts.NewTransactionClient(cfg.TxnHost, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if clientErr != nil {
		log.Fatal(clientErr)
	}
	claimService, claimErr := api.NewContractIntegrationRPCService(cfg.Port, []grpc.ServerOption{grpc.EmptyServerOption{}}, txnClient)
	if claimErr != nil {
		log.Fatal(claimErr)
	}
	log.Fatal(<-claimService.Channel)
}
