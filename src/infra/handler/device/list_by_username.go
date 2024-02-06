package device_handler

import (
	"context"

	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
)

type listDevicesByUsernameHandler struct {
	listDevicesByUsernameUseCase *device_use_case.ListDevicesByUsernameUseCase
}

func NewListDevicesByUsernameHandler(listDevices *device_use_case.ListDevicesByUsernameUseCase) handler.Handler {
	return &listDevicesByUsernameHandler{listDevices}
}

func (h *listDevicesByUsernameHandler) Handle(ctx context.Context, request *handler.HttpRequest) (*handler.HttpResponse, error) {
	username := request.Query("username")
	devices, err := h.listDevicesByUsernameUseCase.Execute(ctx, username)
	if err != nil {
		return nil, err
	}

	return &handler.HttpResponse{
		Status: 200,
		Body:   devices,
	}, nil
}
