package device_handler

import (
	"context"

	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
	"github.com/DTunnel0/CheckUser-Go/src/infra/handler"
)

type countDevicesHandler struct {
	countDevicesUseCase *device_use_case.CountDevicesUseCase
}

func NewCountDevicesHandler(countDevicesUseCase *device_use_case.CountDevicesUseCase) handler.Handler {
	return &countDevicesHandler{countDevicesUseCase}
}

func (h *countDevicesHandler) Handle(ctx context.Context, request *handler.HttpRequest) (*handler.HttpResponse, error) {
	count, err := h.countDevicesUseCase.Execute(ctx)
	if err != nil {
		return nil, err
	}

	return &handler.HttpResponse{
		Status: 200,
		Body:   map[string]int{"count": count},
	}, nil
}
