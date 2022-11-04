package main

import (
	"contract-service/utils"
	"log"
)

func main() {
	liveProbeErr := make(chan string)
	probe := utils.NewProbe()

	probe.Serve(liveProbeErr)


	log.Fatal(<-liveProbeErr)
}

