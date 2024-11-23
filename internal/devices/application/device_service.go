package application

import (
	"time"

	"github.com/mehmetymw/thundio/internal/devices/domain"
)

type DeviceService struct {
	Repository domain.DeviceRepository
}

func NewDeviceService(repo domain.DeviceRepository) *DeviceService {
	return &DeviceService{
		Repository: repo,
	}
}

func (s *DeviceService) RegisterDevice(name, deviceType string) (*domain.Device, error) {
	device := &domain.Device{
		Name:      name,
		Type:      deviceType,
		Status:    domain.Inactive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	deviceID, err := s.Repository.Save(device)
	if err != nil {
		return nil, err
	}

	device.ID = deviceID
	return device, nil
}

func (s *DeviceService) GetDeviceByID(id domain.DeviceID) (*domain.Device, error) {
	device, err := s.Repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (s *DeviceService) ListDevices() ([]*domain.Device, error) {
	devices, err := s.Repository.ListDevices()
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *DeviceService) ListDevicesByStatus(status domain.DeviceStatus) ([]*domain.Device, error) {
	devices, err := s.Repository.ListDevicesByStatus(status)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *DeviceService) UpdateDeviceStatus(id domain.DeviceID, status domain.DeviceStatus) error {
	err := s.Repository.UpdateStatus(id, status)
	if err != nil {
		return err
	}
	return nil
}
