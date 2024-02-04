package device_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type ListDevicesByUsernameUseCase struct {
	deviceRepository contract.DeviceRepository
}

func NewListDevicesByUsernameUseCase(deviceRepository contract.DeviceRepository) *ListDevicesByUsernameUseCase {
	return &ListDevicesByUsernameUseCase{
		deviceRepository: deviceRepository,
	}
}

func (l *ListDevicesByUsernameUseCase) Execute(ctx context.Context, username string) ([]*string, error) {
	devices, err := l.deviceRepository.ListByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	result := make([]*string, len(devices))
	for i, device := range devices {
		result[i] = &device.ID
	}
	return result, nil
}
