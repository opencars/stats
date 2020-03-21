package apiserver

import (
	"log"
	"net/http"

	"github.com/shal/statsd/pkg/config"
	"github.com/shal/statsd/pkg/store/sqlstore"
)

// Start starts http server on the specified addr.
func Start(conf *config.Database, addr string) error {
	store, err := sqlstore.New(conf)
	if err != nil {
		return err
	}

	srv := newServer(store)

	log.Printf("Server is listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
