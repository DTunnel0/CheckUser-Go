package repository

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type systemUserRepository struct {
	userDAO contract.UserDAO
}

func NewSystemUserRepository(userDAO contract.UserDAO) contract.UserRepository {
	return &systemUserRepository{
		userDAO: userDAO,
	}
}

func (r *systemUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := r.userDAO.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
