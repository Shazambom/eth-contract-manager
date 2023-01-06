package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//DO NOT IMPLEMENT WITH gRPC SERVICES, USE grpc_health_probe FOR HEALTH CHECKING OF gRPC SERVICES

type LiveResponse struct {
	Body string `json:"body"`
}

type Probe struct {
	router *gin.Engine
}

func NewProbe() *Probe {
	probe := &Probe{router: gin.Default()}
	probe.router.GET("/", probe.Handle)
	return probe
}

func (p *Probe) Serve(port int, err chan string) {
	go func() {
		err <- p.router.Run(fmt.Sprintf(":%d", port)).Error()
	}()
}

func (p *Probe) Handle(ctx *gin.Context) {
	resp := &LiveResponse{Body: "Service is alive"}
	ctx.JSON(http.StatusOK, resp)
	log.Println(resp.Body)
}
