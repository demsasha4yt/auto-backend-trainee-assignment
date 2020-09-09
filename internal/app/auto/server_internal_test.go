package auto

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

var (
	databaseURL string
)

func TestServer_handleShortenURL(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("links")
	s := newServer(sqlstore.New(db))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "Valid",
			payload: map[string]string{
				"url": "google.com",
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "NO url",
			payload: map[string]string{
				"url2222": "google.com",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name: "Wrong URL",
			payload: map[string]string{
				"url": "google",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
		{
			name:         "Invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/shorten_url", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestServer_handleRedirectBase62(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("links")
	s := newServer(sqlstore.New(db))

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/api/shorten_url", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, tc.expectedCode)
		})
	}
}

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=192.168.99.100 dbname=data user=postgres password=pass sslmode=disable"
	}
	os.Exit(m.Run())
}
