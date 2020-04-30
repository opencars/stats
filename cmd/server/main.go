package main

import (
	"flag"
	"log"
	"os"

	"github.com/shal/statsd/pkg/apiserver"

	"github.com/shal/statsd/pkg/config"
)

func main() {
	var path string

	flag.StringVar(&path, "config", "./config/config.toml", "Path to the configuration file")

	flag.Parse()

	// Get configuration.
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	conf, err := config.New(f)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(&conf.DB, ":8080"); err != nil {
		log.Fatal(err)
	}
}
