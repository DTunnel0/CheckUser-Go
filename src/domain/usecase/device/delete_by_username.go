package device_use_case

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
)

type DeleteDevicesByUsername struct {
	deviceRepository contract.DeviceRepository
}

func NewDeleteDevicesByUsername(deviceRepository contract.DeviceRepository) *DeleteDevicesByUsername {
	return &DeleteDevicesByUsername{
		deviceRepository: deviceRepository,
	}
}

func (d *DeleteDevicesByUsername) Execute(ctx context.Context, username string) error {
	return d.deviceRepository.DeleteByUsername(ctx, username)
}
