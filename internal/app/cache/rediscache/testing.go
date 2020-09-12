package rediscache

import (
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

// TestCache returns test redis cache entity
func TestCache(t *testing.T) (*Cache, func()) {
	t.Helper()
	pool := newPool("redis:6379")
	return New(pool, "AVITO_TEST"), func() {
		pool.Close()
	}
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
