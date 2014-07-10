package main

import (
	"time"

	"github.com/joiggama/keeper/config"
	"github.com/joiggama/keeper/deis"
)

func main() {
	config.Init()

	tick := time.NewTicker(*config.INTERVAL)

	for {
		<-tick.C

		deis.Monitor()
	}
}
