package user_use_case

import (
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"golang.org/x/net/context"
)

type DetailUserOutput struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	ExpiresAt   string `json:"expires_at"`
	ExpiresDays int    `json:"expires_days"`
	Limit       int    `json:"limit"`
	Connections int    `json:"connections"`
}

type DetailUserUseCase struct {
	userRepository  contract.UserRepository
	countConnection contract.CountConnection
}

func NewDetailUserUseCase(
	userRepository contract.UserRepository,
	countConnection contract.CountConnection,
) *DetailUserUseCase {
	return &DetailUserUseCase{
		userRepository:  userRepository,
		countConnection: countConnection,
	}
}

func (c *DetailUserUseCase) Execute(ctx context.Context, username string) (*DetailUserOutput, error) {
	user, err := c.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	connections, err := c.countConnection.ByUsername(ctx, user.Username)
	if err != nil {
		connections = 0
	}

	return &DetailUserOutput{
		ID:          user.ID,
		Username:    user.Username,
		ExpiresAt:   user.ExpiresAt.Format("01/01/2006"),
		Limit:       user.Limit,
		ExpiresDays: int(user.ExpiresAt.Sub(time.Now()).Hours() / 24),
		Connections: connections,
	}, nil
}
