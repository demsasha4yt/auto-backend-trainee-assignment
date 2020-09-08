package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/models"
)

const (
	shortURLEmpty string = ""
)

// LinksRepository ...
type LinksRepository struct {
	store *Store
}

// Create creates a new shorten_link
func (r *LinksRepository) Create(u *models.Links) error {
	if err := u.Validate(); err != nil {
		return err
	}
	if err := r.store.db.QueryRow(
		"INSERT INTO links(shorten_url, original_url) VALUES($1, $2) RETURNING id",
		shortURLEmpty,
		u.URL,
	).Scan(&u.ID); err != nil {
		return err
	}
	u.MakeShorten()
	return nil
}

// FindByID finds URL by ID in db
func (r *LinksRepository) FindByID(id int64) (*models.Links, error) {
	l := &models.Links{}
	if err := r.store.db.QueryRow(
		"SELECT id, original_url FROM links WHERE id = $1",
		id,
	).Scan(
		&l.ID,
		&l.URL,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Record not found")
		}
		return nil, err
	}
	return l, nil
}
