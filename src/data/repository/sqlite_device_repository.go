package repository

import (
	"context"
	"database/sql"

	"github.com/DTunnel0/CheckUser-Go/src/domain/contract"
	"github.com/DTunnel0/CheckUser-Go/src/domain/entity"
	"github.com/labstack/gommon/log"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

const dbURI = "./db.sqlite3"

type SQLiteDeviceRepository struct {
	db *sql.DB
}

func NewSQLiteDeviceRepository() contract.DeviceRepository {
	db, err := sql.Open("sqlite3", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS devices (
		id TEXT PRIMARY KEY,
		username TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	return &SQLiteDeviceRepository{db: db}
}

func (r *SQLiteDeviceRepository) Save(ctx context.Context, device *entity.Device) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO devices (id, username) VALUES (?, ?)", device.ID, device.Username)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLiteDeviceRepository) Exists(ctx context.Context, device *entity.Device) bool {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM devices WHERE id = ?", device.ID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (r *SQLiteDeviceRepository) ListByUsername(ctx context.Context, username string) ([]*entity.Device, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, username FROM devices WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []*entity.Device{}
	for rows.Next() {
		device := &entity.Device{}
		err := rows.Scan(&device.ID, &device.Username)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *SQLiteDeviceRepository) ListAll(ctx context.Context) ([]*entity.Device, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, username FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []*entity.Device{}
	for rows.Next() {
		device := &entity.Device{}
		err := rows.Scan(&device.ID, &device.Username)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (r *SQLiteDeviceRepository) DeleteByUsername(ctx context.Context, username string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM devices WHERE username = ?", username)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLiteDeviceRepository) CountByUsername(ctx context.Context, username string) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM devices WHERE username = ?", username).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func DeleteDB() {
	db, err := sql.Open("sqlite3", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`DROP TABLE IF EXISTS devices`)
	if err != nil {
		log.Fatal(err)
	}
}
