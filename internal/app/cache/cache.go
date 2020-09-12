package cache

import "time"

// Cache interface
type Cache interface {
	Get(key interface{}) (string, error)
	Set(key, val interface{}, ttl time.Duration) error
}
