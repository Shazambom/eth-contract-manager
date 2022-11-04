package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Body string `json:"body"`
}

type Probe struct {
	router *mux.Router
}

func NewProbe() *Probe {
	probe := &Probe{router: mux.NewRouter()}
	probe.router.HandleFunc("/", probe.Handle)
	return probe
}

func (p *Probe) Serve(err chan string) {
	go func() {
		err <- http.ListenAndServe(":8080", p.router).Error()
	}()
}

func (p *Probe) Handle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	resp := Response{Body: "Service is alive"}
	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
	log.Println(resp.Body)
}
