package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/opencars/statsd/pkg/model"
	"github.com/opencars/statsd/pkg/store/sqlstore"

	"github.com/opencars/statsd/pkg/eventapi"

	"github.com/nats-io/nats.go"

	"github.com/opencars/statsd/pkg/config"
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
