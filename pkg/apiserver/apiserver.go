package apiserver

import (
	"log"
	"net/http"

	"github.com/opencars/statsd/pkg/store"
)

// Start starts http server on the specified addr.
func Start(store store.Store, addr string) error {
	srv := newServer(store)

	log.Printf("Server is listening on %s...\n", addr)
	return http.ListenAndServe(addr, srv)
}
