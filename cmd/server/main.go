package main

import (
	"flag"
	"log"

	"github.com/opencars/statsd/pkg/apiserver"
	"github.com/opencars/statsd/pkg/config"
	"github.com/opencars/statsd/pkg/store/sqlstore"
)

func main() {
	var path string

	flag.StringVar(&path, "config", "./config/config.yaml", "Path to the configuration file")

	flag.Parse()

	conf, err := config.New(path)
	if err != nil {
		log.Fatal(err)
	}

	store, err := sqlstore.New(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(store, ":8080"); err != nil {
		log.Fatal(err)
	}
}
