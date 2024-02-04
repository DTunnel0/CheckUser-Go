package presenter

import (
	"context"
	"fmt"
	"log"

	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
)

type ListDevicesPresenter struct {
	listDevicesUseCase *device_use_case.ListDevicesUseCase
}

func NewListDevicesPresenter(listDevicesUseCase *device_use_case.ListDevicesUseCase) *ListDevicesPresenter {
	return &ListDevicesPresenter{
		listDevicesUseCase: listDevicesUseCase,
	}
}
func (p *ListDevicesPresenter) Present(ctx context.Context) {
	devices, err := p.listDevicesUseCase.Execute(ctx)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if len(devices) == 0 {
		fmt.Println("No devices found")
		return
	}

	message := fmt.Sprintf("----------------")
	for _, device := range devices {
		message += fmt.Sprintf("\n%s", *device)
	}
	message += "\n----------------"
	fmt.Println(message)
}
