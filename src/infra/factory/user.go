package factory

import (
	"github.com/DTunnel0/CheckUser-Go/src/data"
	"github.com/DTunnel0/CheckUser-Go/src/data/cache"
	"github.com/DTunnel0/CheckUser-Go/src/data/connection"
	"github.com/DTunnel0/CheckUser-Go/src/data/dao"
	"github.com/DTunnel0/CheckUser-Go/src/data/repository"
	user_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/user"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
	user_handler "github.com/DTunnel0/CheckUser-Go/src/infra/handler/user"
)

func MakeCheckUserHandler() handler.Handler {
	executor := data.NewBashExecutor()
	userCaseService := cache.NewUserCacheService()
	userDAO := dao.NewUserDAO(executor)
	userRepository := repository.NewSystemUserRepository(userDAO, userCaseService)
	deviceRepository := repository.NewSQLiteDeviceRepository()
	checkUserUseCase := user_use_case.NewCheckUserUseCase(userRepository, deviceRepository)
	return user_handler.NewCheckUserHandler(checkUserUseCase)
}

func MakeCountConnectionsHandler() handler.Handler {
	executor := data.NewBashExecutor()
	countSSH := connection.NewSSHConnection(executor)
	countSSH.SetNext(connection.NewOpenVPNConnection(connection.NewAUXOpenVPNConnection("127.0.0.1", 7505)))
	countConnectionCacheService := cache.NewCountConnectionCacheService()
	countConnectionsUseCase := user_use_case.NewCountConnectionsUseCase(countSSH, countConnectionCacheService)
	return user_handler.NewCountConnectionsHandler(countConnectionsUseCase)
}

func MakeDetailsUserHandler() handler.Handler {
	executor := data.NewBashExecutor()
	userCaseService := cache.NewUserCacheService()
	userDAO := dao.NewUserDAO(executor)
	userRepository := repository.NewSystemUserRepository(userDAO, userCaseService)
	countSSH := connection.NewSSHConnection(executor)
	countSSH.SetNext(connection.NewOpenVPNConnection(connection.NewAUXOpenVPNConnection("127.0.0.1", 7505)))
	checkUserUseCase := user_use_case.NewDetailUserUseCase(userRepository, countSSH)
	return user_handler.NewDetailUserHandler(checkUserUseCase)
}
