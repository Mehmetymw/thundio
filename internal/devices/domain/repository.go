package domain

type DeviceRepository interface {
	Save(device *Device) (DeviceID, error)
	GetByID(id DeviceID) (*Device, error)
	UpdateStatus(id DeviceID, status DeviceStatus) error
	ListDevices() ([]*Device, error)
	ListDevicesByStatus(DeviceStatus) ([]*Device, error)
}
