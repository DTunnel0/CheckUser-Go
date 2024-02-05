package cache

import (
	"sync"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type countConnectionCacheService struct {
	cache sync.Map
}

func NewCountConnectionCacheService() contract.CountConnectionCacheService {
	return &countConnectionCacheService{}
}

type countCachedResult struct {
	value    int
	expireAt time.Time
}

func (c *countConnectionCacheService) Set(key string, value int, ttl time.Duration) {
	expireAt := time.Now().Add(ttl)
	c.cache.Store(key, &countCachedResult{
		value:    value,
		expireAt: expireAt,
	})
}

func (c *countConnectionCacheService) Get(key string) (int, bool) {
	result, found := c.cache.Load(key)
	if !found {
		return 0, false
	}

	cached := result.(*countCachedResult)
	if time.Now().After(cached.expireAt) {
		c.cache.Delete(key)
		return 0, false
	}
	return cached.value, true
}
