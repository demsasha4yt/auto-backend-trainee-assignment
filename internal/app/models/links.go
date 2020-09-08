package models

import (
	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/base62"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Links ...
type Links struct {
	ID         int64  `json:"id,omitempty"`
	URL        string `json:"url,omitempty"`
	ShortenURL string `json:"short_url,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}

// Validate links
func (s *Links) Validate() error {
	return validation.ValidateStruct(s,
		validation.Field(&s.URL, validation.Required, is.URL))
}

// MakeShorten makes shorten url and put in s.ShortenURL
func (s *Links) MakeShorten() {
	s.ShortenURL = base62.Encode(s.ID)
}

// MakeID makes ID from encoded string and put in s.ID for finding it latter
func (s *Links) MakeID() error {
	var err error
	s.ID, err = base62.Decode(s.ShortenURL)
	return err
}
