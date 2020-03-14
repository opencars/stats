package sqlstore_test

import (
	"github.com/shal/statsd/pkg/model"
	"github.com/shal/statsd/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
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
