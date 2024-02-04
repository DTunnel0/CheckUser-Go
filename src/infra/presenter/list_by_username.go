package presenter

import (
	"context"
	"fmt"
	"log"

	device_use_case "github.com/DTunnel0/CheckUser-Go/src/domain/usecase/device"
)

type ListDevicesByUsernamePresenter struct {
	listDevicesByUsernameUseCase *device_use_case.ListDevicesByUsernameUseCase
}

func NewListDevicesByUsernamePresenter(listDevicesByUsernameUseCase *device_use_case.ListDevicesByUsernameUseCase) *ListDevicesByUsernamePresenter {
	return &ListDevicesByUsernamePresenter{
		listDevicesByUsernameUseCase: listDevicesByUsernameUseCase,
	}
}

func (p *ListDevicesByUsernamePresenter) Present(ctx context.Context, username string) {
	devices, err := p.listDevicesByUsernameUseCase.Execute(ctx, username)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if len(devices) == 0 {
		log.Printf("No devices found for user %s", username)
		return
	}

	message := fmt.Sprintf("----------[%s]----------", username)
	for _, device := range devices {
		message += fmt.Sprintf("\n%s", *device)
	}
	message += "\n-------------------------"
	fmt.Println(message)
}
