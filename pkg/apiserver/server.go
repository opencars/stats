package apiserver

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"net/http"

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
	headers := handlers.AllowedHeaders([]string{"X-Api-Key", "Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(s.router)
	cors.ServeHTTP(w, r)
}

func (s *server) handleStatsForTimeInterval() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		from, err := strconv.ParseInt(mux.Vars(r)["from"], 10, 64)
		if err != nil {
			// TODO: Return 400.
			return err
		}

		to, err := strconv.ParseInt(mux.Vars(r)["to"], 10, 64)
		if err != nil {
			// TODO: Return 400.
			return err
		}

		activity, err := s.store.Authorization().StatsForPeriod(time.Unix(from, 0), time.Unix(to, 0))
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(activity); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) handleStatsForDuration() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		relative, err := time.ParseDuration(mux.Vars(r)["duration"])
		if err != nil {
			return err
		}

		activity, err := s.store.Authorization().StatsForPeriod(time.Now().Add(-relative), time.Now().Add(time.Hour))
		if err != nil {
			// TODO: Return 400.
			return err
		}

		if err := json.NewEncoder(w).Encode(activity); err != nil {
			return err
		}

		return nil
	}
}

func (s *server) handleActivityByToken() handler.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		token := mux.Vars(r)["token"]

		activity, err := s.store.Authorization().StatsByToken(token)
		if err == sql.ErrNoRows {
			return handler.ErrNotFound
		}

		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(activity); err != nil {
			return err
		}

		return nil
	}
}
