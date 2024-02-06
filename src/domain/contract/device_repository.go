package contract

import (
	"context"

	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
)

type DeviceRepository interface {
	Save(ctx context.Context, device *entity.Device) error
	Exists(ctx context.Context, device *entity.Device) bool
	DeleteByUsername(ctx context.Context, username string) error
	CountByUsername(ctx context.Context, username string) (int, error)
	ListByUsername(ctx context.Context, username string) ([]*entity.Device, error)
	ListAll(ctx context.Context) ([]*entity.Device, error)
	CountAll(ctx context.Context) (int, error)
}
