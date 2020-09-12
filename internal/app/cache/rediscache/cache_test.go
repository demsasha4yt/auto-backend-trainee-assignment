package rediscache_test

import (
	"testing"
	"time"

	"github.com/demsasha4yt/auto-backend-trainee-assignment/internal/app/cache/rediscache"
	"github.com/stretchr/testify/assert"
)

func TestCache_Set(t *testing.T) {
	cache, teardown := rediscache.TestCache(t)
	defer teardown()
	err := cache.Set("testkey", "testval", time.Minute)
	assert.NoError(t, err)
}

func TestCache_Get(t *testing.T) {
	cache, teardown := rediscache.TestCache(t)
	defer teardown()

	err := cache.Set("testkey", "testval", time.Minute)
	assert.NoError(t, err)
	data, err := cache.Get("testkey")
	val := string(data)
	assert.NoError(t, err)
	assert.Equal(t, "testval", val)
}
