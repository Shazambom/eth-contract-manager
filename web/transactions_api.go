package web

import (
	"bitbucket.org/artie_inc/contract-service/contracts"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"fmt"
	"net/http"
)
import "github.com/gin-gonic/gin"


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
	tokens := []*storage.Transaction{}
	for _, txn := range transactions.Transactions {
		token := &storage.Transaction{}
		tokenErr := token.FromRPC(txn)
		if tokenErr != nil {
			ctx.JSON(http.StatusInternalServerError, tokenErr)
		}
		tokens = append(tokens, token)
	}
	ctx.JSON(http.StatusOK, tokens)
}