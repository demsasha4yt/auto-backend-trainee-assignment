package sqlstore_test

import (
	"testing"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/models"
	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestLinksRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("links")

	s := sqlstore.New(db)
	l := models.TestLink(t)

	err := s.Links().Create(l)
	assert.NoError(t, err)
	assert.NotNil(t, l)
}

func TestLinksRepository_FindByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("links")

	s := sqlstore.New(db)
	l := models.TestLink(t)

	err := s.Links().Create(l)
	assert.NoError(t, err)
	assert.NotNil(t, l)
	l2, err := s.Links().FindByID(l.ID)
	assert.NoError(t, err)
	assert.NotNil(t, l2)
	l3, err := s.Links().FindByID(100500)
	assert.Error(t, err)
	assert.Nil(t, l3)
}
