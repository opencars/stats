package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/shal/statsd/pkg/model"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/shal/statsd/internal/config"
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

	conn, err := nats.Connect(conf.EventAPI.Address())
	if err != nil {
		log.Fatal(err)
	}

	events := make(chan *nats.Msg, 64)
	_, err = conn.ChanSubscribe("events.auth.new", events)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening...")
	for event := range events {
		var tmp model.Event
		if err := json.Unmarshal(event.Data, &tmp); err != nil {
			log.Fatal(err)
		}

		fmt.Println(tmp)
	}
}
