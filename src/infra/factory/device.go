package factory

import (
	"github.com/DTunnel0/CheckUser-Go/src/data/repository"
	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
	device_handler "github.com/DTunnel0/CheckUser-Go/src/infra/handler/device"
	"github.com/DTunnel0/CheckUser-Go/src/infra/presenter"
)

func MakeListDevicesPresenter() *presenter.ListDevicesPresenter {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	listDevicesUseCase := device_use_case.NewListDevicesUseCase(deviceRepository)
	listDevicesPresenter := presenter.NewListDevicesPresenter(listDevicesUseCase)
	return listDevicesPresenter
}

func MakeListDevicesByUsernamePresenter() *presenter.ListDevicesByUsernamePresenter {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	listDevicesByUsernameUseCase := device_use_case.NewListDevicesByUsernameUseCase(deviceRepository)
	listDevicesByUsernamePresenter := presenter.NewListDevicesByUsernamePresenter(listDevicesByUsernameUseCase)
	return listDevicesByUsernamePresenter
}

func MakeDeleteDeviceByUsernamePresenter() *presenter.DeleteDevicesPresenter {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	deleteDeviceByUsernameUseCase := device_use_case.NewDeleteDevicesByUsername(deviceRepository)
	deleteDeviceByUsernamePresenter := presenter.NewDeleteDevicesPresenter(deleteDeviceByUsernameUseCase)
	return deleteDeviceByUsernamePresenter
}

func MakeListDevicesHandler() handler.Handler {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	use_case := device_use_case.NewListDevicesUseCase(deviceRepository)
	handler := device_handler.NewListDevicesHandler(use_case)
	return handler
}

func MakeListDevicesByUsernameHandler() handler.Handler {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	use_case := device_use_case.NewListDevicesByUsernameUseCase(deviceRepository)
	handler := device_handler.NewListDevicesByUsernameHandler(use_case)
	return handler
}

func MakeCountDevicesHandler() handler.Handler {
	deviceRepository := repository.NewSQLiteDeviceRepository()
	use_case := device_use_case.NewCountDevicesUseCase(deviceRepository)
	handler := device_handler.NewCountDevicesHandler(use_case)
	return handler
}
