package apiserver

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/opencars/edrmvs/pkg/store"
	"github.com/shal/statsd/pkg/apiserver/handler"
)

func to(r *http.Request) (int64, error) {
	to := mux.Vars(r)["to"]
	if to == "" {
		return time.Now().UTC().Unix(), nil
	}

	result, err := strconv.ParseInt(to, 10, 64)
	if err != nil {
		return 0, handler.ErrInvalidTimestamp
	}

	return result, nil
}

func handleErr(err error) error {
	switch err {
	case store.ErrRecordNotFound:
		return handler.ErrNotFound
	default:
		return err
	}
}
