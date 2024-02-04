package user_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type CountConnectionsUseCase struct {
	connection contract.Connection
}

func NewCountConnectionsUseCase(connection contract.Connection) *CountConnectionsUseCase {
	return &CountConnectionsUseCase{
		connection: connection,
	}
}

func (c *CountConnectionsUseCase) Execute(ctx context.Context) int {
	count := c.connection.Count(ctx)
	return count
}
