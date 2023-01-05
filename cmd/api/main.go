package main

import (
	"contract-service/web"
	"log"
)

func main() {
	liveProbeErr := make(chan string)
	probe := web.NewProbe()

	probe.Serve(8080, liveProbeErr)


	log.Fatal(<-liveProbeErr)
}

