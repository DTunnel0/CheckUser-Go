package cache

import (
	"sync"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type userCacheService struct {
	cache sync.Map
}

type cachedResult struct {
	item     *entity.User
	expireAt time.Time
}

func NewUserCacheService() contract.UserCacheService {
	return &userCacheService{}
}

func (u *userCacheService) Set(value *entity.User, ttl time.Duration) {
	expireAt := time.Now().Add(ttl)
	u.cache.Store(value.Username, &cachedResult{
		item:     value,
		expireAt: expireAt,
	})
}

func (u *userCacheService) Get(username string) (*entity.User, bool) {
	result, found := u.cache.Load(username)
	if !found {
		return nil, false
	}

	cached := result.(*cachedResult)
	if time.Now().After(cached.expireAt) {
		u.cache.Delete(username)
		return nil, false
	}
	return cached.item, true
}
