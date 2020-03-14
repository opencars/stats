package sqlstore

import (
	"github.com/shal/statsd/pkg/model"
)

type AuthRepository struct {
	store *Store
}

func (r *AuthRepository) Create(auth *model.Authorization) error {
	row := r.store.db.QueryRow(`
		INSERT INTO authorizations (
			enabled, token, error, ip, name, status, timestamp
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		) RETURNING id, created_at`,
		auth.Enabled, auth.Token, auth.Error,
		auth.IP, auth.Name, auth.Status, auth.Time,
	)

	if err := row.Scan(&auth.ID, &auth.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) FindByID(id int64) (*model.Authorization, error) {
	var auth model.Authorization

	err := r.store.db.Get(&auth,
		`SELECT id, enabled, token, error, ip, name, status, timestamp, created_at
	   FROM authorizations
	   WHERE id = $1`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &auth, nil
}
