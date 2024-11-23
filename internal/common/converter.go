package common

import (
	"errors"

	"github.com/mehmetymw/thundio/internal/devices/db/generated"
	"github.com/mehmetymw/thundio/internal/devices/domain"
)

func DomainConvertDbToDomain(genDevice generated.Device) (*domain.Device, error) {
	if genDevice.ID < 0 {
		return nil, errors.New("invalid id")
	}

	status := domain.DeviceStatus(genDevice.Status)
	if !status.IsValid() {
		return nil, errors.New("invalid status")
	}

	return &domain.Device{
		ID:        domain.DeviceID(genDevice.ID),
		Name:      genDevice.Name,
		Type:      genDevice.Type,
		Status:    status,
		CreatedAt: genDevice.CreatedAt,
		UpdatedAt: genDevice.UpdatedAt,
	}, nil
}

func DomainConvertDomainToDb(domainDevice domain.Device) (*generated.Device, error) {

	status := generated.DeviceStatus(domainDevice.Status)
	if !domain.DeviceStatus.IsValid(domain.DeviceStatus(status)) {
		return nil, errors.New("invalid status")
	}

	return &generated.Device{
		ID:        int32(domainDevice.ID),
		Name:      domainDevice.Name,
		Type:      domainDevice.Type,
		Status:    status,
		CreatedAt: domainDevice.CreatedAt,
		UpdatedAt: domainDevice.UpdatedAt,
	}, nil
}
