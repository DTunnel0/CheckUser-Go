package device_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type ListDevicesUseCase struct {
	deviceRepository contract.DeviceRepository
}

func NewListDevicesUseCase(deviceRepository contract.DeviceRepository) *ListDevicesUseCase {
	return &ListDevicesUseCase{
		deviceRepository: deviceRepository,
	}
}

func (l *ListDevicesUseCase) Execute(ctx context.Context) ([]*string, error) {
	devices, err := l.deviceRepository.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*string, len(devices))
	for i, device := range devices {
		result[i] = &device.ID
	}
	return result, nil
}
