package user_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type CountConnectionsByUsernameUseCase struct {
	connection contract.Connection
}

func NewCountConnectionsByUsernameUseCase(connection contract.Connection) *CountConnectionsByUsernameUseCase {
	return &CountConnectionsByUsernameUseCase{
		connection: connection,
	}
}

func (c *CountConnectionsByUsernameUseCase) Execute(ctx context.Context, username string) (int, error) {
	count, err := c.connection.CountByUsername(ctx, username)
	if err != nil {
		return 0, err
	}
	return count, nil
}
