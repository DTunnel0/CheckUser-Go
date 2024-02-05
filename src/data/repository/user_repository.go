package repository

import (
	"context"
	"time"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type systemUserRepository struct {
	userDAO          contract.UserDAO
	userCacheService contract.UserCacheService
}

func NewSystemUserRepository(userDAO contract.UserDAO, userCacheService contract.UserCacheService) contract.UserRepository {
	return &systemUserRepository{
		userDAO:          userDAO,
		userCacheService: userCacheService,
	}
}

func (r *systemUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, found := r.userCacheService.Get(username)
	if found {
		return user, nil
	}

	user, err := r.userDAO.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	r.userCacheService.Set(user, time.Minute*30)
	return user, nil
}
