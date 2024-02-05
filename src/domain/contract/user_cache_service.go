package contract

import (
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type UserCacheService interface {
	Set(value *entity.User, ttl time.Duration)
	Get(username string) (*entity.User, bool)
}
