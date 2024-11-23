package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mehmetymw/thundio/internal/devices/db/generated"
	"github.com/mehmetymw/thundio/internal/devices/domain"
)

type DeviceRepository struct {
	q *generated.Queries
}

func NewDeviceRepository(db *sql.DB) *DeviceRepository {
	return &DeviceRepository{
		q: generated.New(db),
	}
}

func (r *DeviceRepository) Save(device *domain.Device) (*domain.DeviceID, error) {
	arg := generated.CreateDeviceParams{
		Name:      device.Name,
		Type:      device.Type,
		Status:    generated.DeviceStatus(device.Status),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}
	id, err := r.q.CreateDevice(context.Background(), arg)
	if err != nil {
		return nil, fmt.Errorf("failed to save device: %w", err)
	}

	deviceID := domain.DeviceID(id)
	return &deviceID, nil
}

func (r *DeviceRepository) GetByID(id domain.DeviceID) (*domain.Device, error) {
	genDevice, err := r.q.GetDeviceByID(context.Background(), int32(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get device by ID %d: %w", id, err)
	}

	// Dönüşüm işlemi burada yapılır
	device := &domain.Device{
		ID:        domain.DeviceID(genDevice.ID),
		Name:      genDevice.Name,
		Type:      genDevice.Type,
		Status:    domain.DeviceStatus(genDevice.Status),
		CreatedAt: genDevice.CreatedAt,
		UpdatedAt: genDevice.UpdatedAt,
	}

	return device, nil
}

func (r *DeviceRepository) UpdateStatus(id domain.DeviceID, status domain.DeviceStatus) error {
	arg := generated.UpdateDeviceStatusParams{
		ID:        int32(id),
		Status:    generated.DeviceStatus(status),
		UpdatedAt: time.Now(),
	}
	if err := r.q.UpdateDeviceStatus(context.Background(), arg); err != nil {
		return fmt.Errorf("failed to update device status for ID %d: %w", id, err)
	}
	return nil
}

func (r *DeviceRepository) ListDevices() ([]*domain.Device, error) {
	genDevices, err := r.q.ListDevices(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no devices found")
		}
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}

	// Dönüşüm işlemi burada yapılır
	devices := make([]*domain.Device, len(genDevices))
	for i, row := range genDevices {
		devices[i] = &domain.Device{
			ID:        domain.DeviceID(row.ID),
			Name:      row.Name,
			Type:      row.Type,
			Status:    domain.DeviceStatus(row.Status),
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}

	return devices, nil
}

func (r *DeviceRepository) ListDevicesByStatus(status domain.DeviceStatus) ([]*domain.Device, error) {
	genDevices, err := r.q.ListDevicesByStatus(context.Background(), generated.DeviceStatus(status))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no devices found for status %s", status)
		}
		return nil, fmt.Errorf("failed to list devices by status %s: %w", status, err)
	}

	// Dönüşüm işlemi burada yapılır
	devices := make([]*domain.Device, len(genDevices))
	for i, row := range genDevices {
		devices[i] = &domain.Device{
			ID:        domain.DeviceID(row.ID),
			Name:      row.Name,
			Type:      row.Type,
			Status:    domain.DeviceStatus(row.Status),
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}

	return devices, nil
}
