package rediscache

import (
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

// Cache ...
type Cache struct {
	password string
	pool     *redigo.Pool
}

// New redis cache
func New(pool *redigo.Pool, password string) *Cache {
	return &Cache{
		password: password,
		pool:     pool,
	}
}

func (s *Cache) getClient() (redigo.Conn, error) {
	c := s.pool.Get()
	if _, err := c.Do("AUTH", s.password); err != nil {
		c.Close()
		return nil, err
	}
	return c, nil
}

// Get ...
func (s *Cache) Get(key interface{}) (string, error) {
	c, err := s.getClient()
	if err != nil {
		return "", err
	}
	defer c.Close()
	data, err := redigo.String(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return data, nil
}

// Set ...
func (s *Cache) Set(key, val interface{}, ttl time.Duration) error {
	c, err := s.getClient()
	if err != nil {
		return err
	}
	defer c.Close()
	_, err = c.Do("SET", key, val)
	if err != nil {
		v := val.(string)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, v, err)
	}
	return nil
}
