package factory

import (
	"github.com/DTunnel0/CheckUser-Go/src/data"
	"github.com/DTunnel0/CheckUser-Go/src/data/connection"
	"github.com/DTunnel0/CheckUser-Go/src/data/repository"
	user_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/user"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
	user_handler "github.com/DTunnel0/CheckUser-Go/src/infra/handler/user"
)

func MakeCheckUserHandler() handler.Handler {
	executor := data.NewBashExecutorWithCache()
	userRepository := repository.NewSystemUserRepository(executor)
	deviceRepository := repository.NewSQLiteDeviceRepository()
	checkUserUseCase := user_use_case.NewCheckUserUseCase(userRepository, deviceRepository)
	return user_handler.NewCheckUserHandler(checkUserUseCase)
}

func MakeCountConnectionsHandler() handler.Handler {
	executor := data.NewBashExecutor()
	ssh := connection.NewSSHConnection(executor)
	ssh.SetNext(connection.NewOpenVPNConnection(connection.NewAUXOpenVPNConnection("127.0.0.1", 7505)))
	countConnectionsUseCase := user_use_case.NewCountConnectionsUseCase(ssh)
	return user_handler.NewCountConnectionsHandler(countConnectionsUseCase)
}
