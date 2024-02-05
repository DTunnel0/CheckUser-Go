package contract

import "time"

type CountConnectionCacheService interface {
	Set(key string, value int, ttl time.Duration)
	Get(key string) (int, bool)
}
