package auto

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/cache/rediscache"
	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/store/sqlstore"
	"github.com/gomodule/redigo/redis"
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

	pool := newPool("redis:6379")
	cache := rediscache.New(pool, "AVITO_TEST")

	server := newServer(store, cache)
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

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
