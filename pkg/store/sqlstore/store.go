package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/statsd/pkg/config"
	"github.com/opencars/statsd/pkg/store"
)

// Store is an implementation of store.Store interface based on SQL.
type Store struct {
	db             *sqlx.DB
	authRepository *AuthRepository
}

// Authorization returns repository, which is responsible for authorizations.
func (s *Store) Authorization() store.AuthRepository {
	if s.authRepository != nil {
		return s.authRepository
	}

	s.authRepository = &AuthRepository{
		store: s,
	}

	return s.authRepository
}

// New returns new instance of Store.
func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Database,
		settings.SSLMode,
		settings.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
