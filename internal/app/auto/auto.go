package auto

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store/sqlstore"
	_ "github.com/lib/pq" // ...
)

// Start application
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	server := newServer(store)
	http.ListenAndServe(config.BindAddr, server)
	return nil
}

func newDB(databaseURL string) (*sql.DB, error) {
	log.Println("[SQL]: Connecting to SQL..")
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}
	log.Println("[SQL]: Success!")
	return db, nil
}
