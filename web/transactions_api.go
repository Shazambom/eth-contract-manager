package web

import (
	"contract-service/contracts"
	pb "contract-service/proto"
	"contract-service/storage"
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
	tokens := []*storage.Token{}
	for _, txn := range transactions.Transactions {
		token := &storage.Token{}
		tokenErr := token.FromRPC(txn)
		if tokenErr != nil {
			ctx.JSON(http.StatusInternalServerError, tokenErr)
		}
		tokens = append(tokens, token)
	}
	ctx.JSON(http.StatusOK, tokens)
}