package store

import (
	"time"

	"github.com/shal/statsd/pkg/model"
)

type AuthRepository interface {
	Create(auth *model.Authorization) error
	FindByID(id int64) (*model.Authorization, error)
	StatsForPeriod(from, to time.Time) ([]model.AuthStat, error)
	StatsByToken(token string) (*model.TokenStat, error)
	StatsByTokenPeriod(from, to time.Time, token string) (*model.TokenStat, error)
}
