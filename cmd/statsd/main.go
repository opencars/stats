package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shal/statsd/pkg/model"
	"github.com/shal/statsd/pkg/store/sqlstore"

	"github.com/shal/statsd/pkg/eventapi"

	"github.com/nats-io/nats.go"
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

	store, err := sqlstore.New(&conf.DB)
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
		var tmp eventapi.Event
		if err := json.Unmarshal(event.Data, &tmp); err != nil {
			log.Fatal(err)
		}

		if tmp.Kind == eventapi.EventAuthorizationKind {
			var auth model.Authorization

			if err := json.Unmarshal(tmp.Data, &auth); err != nil {
				log.Fatal(err)
			}

			if err := store.Authorization().Create(&auth); err != nil {
				log.Fatal(err)
			}

			continue
		}

		log.Printf("[WARNING] %v\n", tmp)
	}
}