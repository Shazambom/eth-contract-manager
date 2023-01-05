package web

import "contract-service/contracts"
import "github.com/gin-gonic/gin"

type TransactionAPI struct {
	client contracts.TransactionClient
	router gin.IRouter
}


