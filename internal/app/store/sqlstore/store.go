package sqlstore

import (
	"database/sql"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store"
)

// Store sqlstore structure
type Store struct {
	db              *sql.DB
	linksRepository *LinksRepository
}

// New creates new sqlstore
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Links return links repository
func (s *Store) Links() store.LinksRepository {
	if s.linksRepository != nil {
		return s.linksRepository
	}

	s.linksRepository = &LinksRepository{
		store: s,
	}

	return s.linksRepository
}
