package web

import (
	"contract-service/contracts"
	pb "contract-service/proto"
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"

//TODO DECIDE: Implement interface pattern for the APIs? Or is it specific enough as an implementation that it doesn't require interfaces.

type TransactionAPI struct {
	client *contracts.TransactionClient
	router *gin.Engine
}

func NewTransactionAPI(client *contracts.TransactionClient) *TransactionAPI {
	return &TransactionAPI{
		client: client,
		router: gin.Default(),
	}
}

func (t *TransactionAPI) Serve(port int, err chan string) {
	t.router.GET(
		"/transactions/:address",
		t.GetTransactionsFromAddress,
		)

	go func() {
		err <- t.router.Run(fmt.Sprintf(":%d", port)).Error()
	}()
}

func (t *TransactionAPI) GetTransactionsFromAddress(ctx *gin.Context) {
	transactions, err := t.client.Client.GetTransactions(ctx, &pb.Address{Address: ctx.Param("address")})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
	}
	//TODO determine if this is the format we want to return things in? I think probably not given the web api will be using this. Cool that it works tho.
	ctx.ProtoBuf(http.StatusOK, transactions)
}