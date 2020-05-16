package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/shal/statsd/pkg/apiserver/handler"
	"github.com/shal/statsd/pkg/store"
)

type TimePeriod string

const (
	MonthTimePeriod TimePeriod = "month"
	DayTimePeriod   TimePeriod = "day"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	srv := server{
		router: mux.NewRouter(),
		store:  store,
	}

	srv.configureRouter()

	return &srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"X-Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) handleActivity() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			return handler.ErrApiKey
		}

		stats, err := s.store.Authorization().StatsByToken(apiKey)
		if err != nil {
			return handleErr(err)
		}

		if err := json.NewEncoder(w).Encode(stats); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) handleActivityPeriod() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		from, err := strconv.ParseInt(mux.Vars(r)["from"], 10, 64)
		if err != nil {
			return handler.ErrInvalidTimestamp
		}

		to, err := to(r)
		if err != nil {
			return err
		}

		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			return handler.ErrApiKey
		}

		stats, err := s.store.Authorization().StatsByTokenPeriod(time.Unix(from, 0), time.Unix(to, 0), apiKey)

		if err != nil {
			return handleErr(err)
		}

		if err := json.NewEncoder(w).Encode(stats); err != nil {
			return err
		}

		return nil
	}
}
