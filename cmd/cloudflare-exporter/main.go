package main

import (
	"cloudflare-status/internal/core"
	"flag"
	"log"
	"os"
)

const VERSION = "v1.2.4"

func main() {
	//go metrics.RecordMetrics()

	configPath := flag.String("config", "../../config/config.yml", "full path to configuration file")
	flag.Parse()

	core, err := core.NewCloudFlareExp(*configPath, VERSION)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	core.Run()
}
