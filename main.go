package main

import (
	"time"

	"github.com/joiggama/keeper/config"
	"github.com/joiggama/keeper/fleetctl"
)

func main() {
	config.Init()

	tick := time.NewTicker(*config.INTERVAL)

	for {
		<-tick.C

		go fleetctl.Monitor()
	}
}
