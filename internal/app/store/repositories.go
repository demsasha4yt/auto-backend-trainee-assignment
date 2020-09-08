package store

import "github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/models"

// LinksRepository interface
type LinksRepository interface {
	Create(*models.Links) error
	FindByID(int64) (*models.Links, error)
}
