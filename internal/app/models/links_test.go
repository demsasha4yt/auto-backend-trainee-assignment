package models_test

import (
	strings "strings"
	"testing"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestLinks_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		l       func() *models.Links
		isValid bool
	}{
		{
			name: "valid",
			l: func() *models.Links {
				return models.TestLink(t)
			},
			isValid: true,
		},
		{
			name: "valid",
			l: func() *models.Links {
				l := models.TestLink(t)
				l.URL = "google"
				return l
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.l().Validate())
			} else {
				assert.Error(t, tc.l().Validate())
			}
		})
	}
}

func TestLinks_AppendHTTP(t *testing.T) {
	l := models.TestLink(t)
	l.URL = "google.com"
	l2 := models.TestLink(t)
	l2.URL = "https://google.com"

	l.AppendHTTP()
	l2.AppendHTTP()
	assert.Equal(t, l.URL, "http://google.com")
	assert.Equal(t, l2.URL, "https://google.com")
}

func TestLinks_MakeShorten_MakeID(t *testing.T) {
	var id int64 = 100500
	base := "q8Y"

	l := models.TestLink(t)
	l.ID = id
	l.MakeShorten()
	assert.Equal(t, l.ShortenURL, base)
	assert.NoError(t, l.MakeID())
	assert.Equal(t, l.ID, id)
}

func TestLinks_PostProcessing(t *testing.T) {
	hostname := "hostname.ru/"
	l := models.TestLink(t)
	l.URL = "https://google.com"
	l.ShortenURL = "q8Y"
	l.PostProcessing(hostname)
	assert.True(t, strings.HasPrefix(l.ShortenURL, hostname))
	assert.True(t, strings.HasSuffix(l.ShortenURL, "q8Y"))
	assert.Equal(t, "hostname.ru/q8Y", l.ShortenURL)
}
