package device_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type ListDevicesOutput struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type ListDevicesUseCase struct {
	deviceRepository contract.DeviceRepository
}

func NewListDevicesUseCase(deviceRepository contract.DeviceRepository) *ListDevicesUseCase {
	return &ListDevicesUseCase{
		deviceRepository: deviceRepository,
	}
}

func (l *ListDevicesUseCase) Execute(ctx context.Context) ([]*ListDevicesOutput, error) {
	devices, err := l.deviceRepository.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*ListDevicesOutput, len(devices))
	for i, device := range devices {
		result[i] = &ListDevicesOutput{
			ID:       device.ID,
			Username: device.Username,
		}
	}
	return result, nil
}
