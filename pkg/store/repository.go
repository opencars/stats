package store

import (
	"github.com/shal/statsd/pkg/model"
)

type AuthRepository interface {
	Create(auth *model.Authorization) error
	FindByID(id int64) (*model.Authorization, error)
}
