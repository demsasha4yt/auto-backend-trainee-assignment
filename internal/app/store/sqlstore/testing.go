package sqlstore

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq" // ...
)

// TestDB ...
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table))
			}
		}
		db.Close()
	}
}
