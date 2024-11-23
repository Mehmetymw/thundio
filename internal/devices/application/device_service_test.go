package application

import (
	"testing"
	"time"

	"github.com/mehmetymw/thundio/internal/devices/domain"
	"github.com/mehmetymw/thundio/internal/devices/infrastructure/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestSaveDevice: Cihaz kaydetme test fonksiyonu
func TestSaveDevice(t *testing.T) {
	mockRepo := new(repository.MockDeviceRepository)

	// Mock repo'nun "Save" metoduna doğru parametreyi vererek doğru şekilde döndürüyoruz
	mockRepo.On("Save", mock.MatchedBy(func(device *domain.Device) bool {
		return device.Name == "Test Device" && device.Type == "Sensor" && device.Status == domain.Inactive
	})).Return(domain.DeviceID(1), nil)

	// DeviceService'i oluşturuyoruz
	service := NewDeviceService(mockRepo)

	// Cihazı kaydediyoruz
	device, err := service.RegisterDevice("Test Device", "Sensor")

	// Test assert
	assert.NoError(t, err)
	assert.NotNil(t, device)
	assert.Equal(t, domain.DeviceID(1), device.ID)

	// Mock beklentilerini doğruluyoruz
	mockRepo.AssertExpectations(t)
}

// TestGetDeviceByID: Cihaz ID'ye göre alma test fonksiyonu
func TestGetDeviceByID(t *testing.T) {
	mockRepo := new(repository.MockDeviceRepository)

	// Mock repo'nun "GetByID" metoduna doğru parametreyi veriyoruz
	mockRepo.On("GetByID", domain.DeviceID(1)).Return(&domain.Device{
		ID:        domain.DeviceID(1),
		Name:      "Test Device",
		Type:      "Sensor",
		Status:    domain.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	// DeviceService'i oluşturuyoruz
	service := NewDeviceService(mockRepo)

	// Cihazı ID'ye göre alıyoruz
	device, err := service.GetDeviceByID(domain.DeviceID(1))

	// Test assert
	assert.NoError(t, err)
	assert.NotNil(t, device)
	assert.Equal(t, "Test Device", device.Name)

	// Mock beklentilerini doğruluyoruz
	mockRepo.AssertExpectations(t)
}

// TestListDevices: Tüm cihazları listeleme test fonksiyonu
func TestListDevices(t *testing.T) {
	mockRepo := new(repository.MockDeviceRepository)

	// Mock repo'nun "ListDevices" metoduna cihazlar dönmesini sağlıyoruz
	mockDevice1 := &domain.Device{
		ID:        domain.DeviceID(1),
		Name:      "Test Device 1",
		Type:      "Sensor",
		Status:    domain.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockDevice2 := &domain.Device{
		ID:        domain.DeviceID(2),
		Name:      "Test Device 2",
		Type:      "Actuator",
		Status:    domain.Inactive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("ListDevices").Return([]*domain.Device{mockDevice1, mockDevice2}, nil)

	// DeviceService'i oluşturuyoruz
	service := NewDeviceService(mockRepo)

	// Cihazları listele
	devices, err := service.ListDevices()

	// Test assert
	assert.NoError(t, err)
	assert.Equal(t, 2, len(devices))
	assert.Equal(t, "Test Device 1", devices[0].Name)

	// Mock beklentilerini doğruluyoruz
	mockRepo.AssertExpectations(t)
}

// TestUpdateDeviceStatus: Cihaz durumunu güncelleme test fonksiyonu
func TestUpdateDeviceStatus(t *testing.T) {
	mockRepo := new(repository.MockDeviceRepository)

	// Mock repo'nun "UpdateStatus" metoduna doğru parametreyi veriyoruz
	mockRepo.On("UpdateStatus", domain.DeviceID(1), domain.Active).Return(nil)

	// DeviceService'i oluşturuyoruz
	service := NewDeviceService(mockRepo)

	// Durum güncelleme işlemi
	err := service.UpdateDeviceStatus(domain.DeviceID(1), domain.Active)

	// Test assert
	assert.NoError(t, err)

	// Mock beklentilerini doğruluyoruz
	mockRepo.AssertExpectations(t)
}
