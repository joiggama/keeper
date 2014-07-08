package main

import (
	"time"

	"github.com/joiggama/unit-keeper/config"
	"github.com/joiggama/unit-keeper/fleetctl"
)

func main() {
	config.Init()

	tick := time.NewTicker(*config.INTERVAL)

	for {
		<-tick.C

		go fleetctl.Monitor()
	}
}
