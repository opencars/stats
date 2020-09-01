package sqlstore_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/statsd/pkg/model"
	"github.com/opencars/statsd/pkg/store/sqlstore"
)

func TestAuthRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("authorizations")

	auth := model.TestAuthorization(t)
	assert.NoError(t, s.Authorization().Create(auth))
	assert.NotNil(t, auth)
	assert.EqualValues(t, 1, auth.ID)
}

func TestAuthRepository_FindByID(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("authorizations")

	auth := model.TestAuthorization(t)
	assert.NoError(t, s.Authorization().Create(auth))
	assert.NotNil(t, auth)
	assert.EqualValues(t, 1, auth.ID)

	actual, err := s.Authorization().FindByID(auth.ID)
	assert.NoError(t, err)
	assert.NotNil(t, actual)
}

func TestAuthRepository_StatsForPeriod(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("authorizations")

	auth := model.TestAuthorization(t)
	assert.NoError(t, s.Authorization().Create(auth))
	assert.NotNil(t, auth)
	assert.EqualValues(t, 1, auth.ID)

	assert.NoError(t, s.Authorization().Create(auth))
	assert.NotNil(t, auth)
	assert.EqualValues(t, 2, auth.ID)

	actual, err := s.Authorization().StatsForPeriod(
		auth.Time.Add(-time.Second), auth.Time.Add(time.Second),
	)
	assert.NoError(t, err)
	assert.Len(t, actual, 1)
}

func TestAuthRepository_StatsByToken(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("authorizations")

	auth1 := model.TestAuthorization(t)
	assert.NoError(t, s.Authorization().Create(auth1))
	assert.NotNil(t, auth1)
	assert.EqualValues(t, 1, auth1.ID)

	auth2 := model.TestAuthorization(t)
	auth2.IP = "127.0.0.2"
	assert.NoError(t, s.Authorization().Create(auth2))
	assert.NotNil(t, auth2)
	assert.EqualValues(t, 2, auth2.ID)

	actual, err := s.Authorization().StatsByToken(auth1.Token)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, actual.Succeed)
	assert.EqualValues(t, 2, actual.Total)
	assert.EqualValues(t, 0, actual.Failed)
}
