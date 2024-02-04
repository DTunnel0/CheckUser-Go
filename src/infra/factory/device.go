package factory

import (
	"github.com/DTunnel0/CheckUser-Go/src/data/repository"
	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
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
