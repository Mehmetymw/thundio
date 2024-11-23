package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	conveter "github.com/mehmetymw/thundio/internal/common"
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
		return nil, fmt.Errorf("device cannot save id:"+string(id), err)
	}

	return (*domain.DeviceID)(&id), nil
}

func (r *DeviceRepository) GetByID(id domain.DeviceID) (*domain.Device, error) {
	genDevice, err := r.q.GetDeviceByID(context.Background(), int32(id))
	if err != nil {
		return nil, fmt.Errorf("device cannot get by id:"+string(id), err)
	}

	device, err := conveter.DomainConvertDbToDomain(genDevice)
	if err != nil {
		return nil, fmt.Errorf("device cannot convert id:"+string(id), err)
	}

	return device, nil

}
func (r *DeviceRepository) UpdateStatus(id domain.DeviceID, status domain.DeviceStatus) error {
	arg := generated.UpdateDeviceStatusParams{
		ID:        int32(id),
		Status:    generated.DeviceStatus(status),
		UpdatedAt: time.Now(),
	}
	return r.q.UpdateDeviceStatus(context.Background(), arg)

}
func (r *DeviceRepository) ListDevices() ([]*domain.Device, error) {
	genDevices, err := r.q.ListDevices(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no devices found")
		}
		return nil, fmt.Errorf("cannot list devices")
	}

	var devices []*domain.Device
	for _, genDevice := range genDevices {
		device, err := conveter.DomainConvertDbToDomain(genDevice)
		if err != nil {
			continue
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (r *DeviceRepository) ListDevicesByStatus(status domain.DeviceStatus) ([]*domain.Device, error) {
	genDevices, err := r.q.ListDevicesByStatus(context.Background(), generated.DeviceStatus(status))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no devices found")
		}
		return nil, fmt.Errorf("cannot list devices")
	}

	var devices []*domain.Device
	for _, genDevice := range genDevices {
		device, err := conveter.DomainConvertDbToDomain(genDevice)
		if err != nil {
			continue
		}

		devices = append(devices, device)
	}

	return devices, nil
}
