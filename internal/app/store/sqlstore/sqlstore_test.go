package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=192.168.99.100 dbname=data user=postgres password=pass sslmode=disable"
	}
	os.Exit(m.Run())
}
