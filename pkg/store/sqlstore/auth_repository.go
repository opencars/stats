package sqlstore

import (
	"context"
	"database/sql"
	"time"

	"github.com/opencars/edrmvs/pkg/store"
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
		auth.IP, auth.Name, auth.Status, auth.Time.UTC(),
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

func (r *AuthRepository) StatsForPeriod(from, to time.Time) ([]model.AuthStat, error) {
	stats := make([]model.AuthStat, 0)

	err := r.store.db.Select(&stats,
		`SELECT token, name, count(*) as amount
		FROM authorizations
		WHERE extract(epoch from timestamp) >= $1 AND
			  extract(epoch from timestamp) <= $2 AND
			  token is not NULL
		GROUP BY token, name
		ORDER BY count(*) DESC`,
		from.UTC().Unix(), to.UTC().Unix(),
	)

	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *AuthRepository) StatsByToken(token string) (*model.TokenStat, error) {
	var stat model.TokenStat

	tx, err := r.store.db.BeginTxx(context.Background(), &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as failed FROM authorizations
		WHERE token = $1 AND status = 'failed'`,
		token,
	)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as succeed FROM authorizations
		WHERE token = $1 AND status = 'succeed'`,
		token,
	)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as total FROM authorizations
		WHERE token = $1`,
		token,
	)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	if stat.Total == 0 {
		return nil, store.ErrRecordNotFound
	}

	return &stat, nil
}

func (r *AuthRepository) StatsByTokenPeriod(from, to time.Time, token string) (*model.TokenStat, error) {
	var stat model.TokenStat

	tx, err := r.store.db.BeginTxx(context.Background(), &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as failed FROM authorizations
		WHERE extract(epoch from timestamp) >= $1 AND
			  extract(epoch from timestamp) <= $2 AND
			  token = $3 AND status = 'failed'`,
		from.UTC().Unix(), to.UTC().Unix(), token,
	)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as succeed FROM authorizations
		WHERE extract(epoch from timestamp) >= $1 AND
			  extract(epoch from timestamp) <= $2 AND
			  token = $3 AND status = 'succeed'`,
		from.UTC().Unix(), to.UTC().Unix(), token,
	)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = tx.Get(&stat,
		`SELECT count(*) as total FROM authorizations
		WHERE extract(epoch from timestamp) >= $1 AND
			  extract(epoch from timestamp) <= $2 AND
			  token = $3`,
		from.UTC().Unix(), to.UTC().Unix(), token,
	)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	if stat.Total == 0 {
		return nil, store.ErrRecordNotFound
	}

	return &stat, nil
}
