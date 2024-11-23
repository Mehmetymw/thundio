package repository

import (
	"time"

	"github.com/mehmetymw/thundio/internal/devices/domain"

	"github.com/stretchr/testify/mock"
)

// MockDeviceRepository: DeviceRepository'nin mock versiyonu
type MockDeviceRepository struct {
	mock.Mock
}

// Save: Cihazı kaydeder
func (m *MockDeviceRepository) Save(device *domain.Device) (domain.DeviceID, error) {
	args := m.Called(device)
	return args.Get(0).(domain.DeviceID), args.Error(1)
}

// GetByID: ID'ye göre cihazı döner
func (m *MockDeviceRepository) GetByID(id domain.DeviceID) (*domain.Device, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Device), args.Error(1)
}

// UpdateStatus: Cihazın durumunu günceller
func (m *MockDeviceRepository) UpdateStatus(id domain.DeviceID, status domain.DeviceStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// ListDevices: Tüm cihazları döner
func (m *MockDeviceRepository) ListDevices() ([]*domain.Device, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Device), args.Error(1)
}

// ListDevicesByStatus: Duruma göre cihazları döner
func (m *MockDeviceRepository) ListDevicesByStatus(status domain.DeviceStatus) ([]*domain.Device, error) {
	args := m.Called(status)
	return args.Get(0).([]*domain.Device), args.Error(1)
}

// NewMockDevice: Yeni bir cihaz oluşturur (Testler için ortak fonksiyon)
func NewMockDevice(id domain.DeviceID, name, deviceType string, status domain.DeviceStatus) *domain.Device {
	return &domain.Device{
		ID:        id,
		Name:      name,
		Type:      deviceType,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
