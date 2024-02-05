package user_use_case

import (
	"context"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type CountConnectionsUseCase struct {
	countConnection      contract.CountConnection
	countConnectionCache contract.CountConnectionCacheService
}

func NewCountConnectionsUseCase(
	countConnection contract.CountConnection,
	countConnectionCache contract.CountConnectionCacheService,
) *CountConnectionsUseCase {
	return &CountConnectionsUseCase{
		countConnection:      countConnection,
		countConnectionCache: countConnectionCache,
	}
}

func (c *CountConnectionsUseCase) Execute(ctx context.Context) (int, error) {
	value, found := c.countConnectionCache.Get("__all__")
	if found {
		return value, nil
	}

	value, err := c.countConnection.All(ctx)
	if err != nil {
		return 0, err
	}

	c.countConnectionCache.Set("__all__", value, time.Second*10)
	return value, nil
}
