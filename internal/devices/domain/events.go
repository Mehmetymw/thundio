package domain

type DeviceRegistered struct {
	DeviceID DeviceID `json:"device_id"`
	Name     string   `json:"string"`
	Type     string   `json:"type"`
}
