package domain

type DeviceStatus string

const (
	Active   DeviceStatus = "Active"
	Inactive DeviceStatus = "Inactive"
)

func (ds DeviceStatus) IsValid() bool {
	switch ds {
	case Active, Inactive:
		return true
	default:
		return false
	}
}
