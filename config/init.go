package config

import (
	"flag"
	"time"
)

var (
	INTERVAL = flag.Duration("interval", 5*time.Second, "Interval between updates (default: 5s)")
)

func Init() {
	flag.Parse()
}
