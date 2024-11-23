package domain

import "time"

type DeviceID int32

type Device struct {
	ID        DeviceID     `json:"id"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Status    DeviceStatus `json:"Status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (d *Device) Activate() {
	d.Status = Active
}

func (d *Device) Deactivate() {
	d.Status = Inactive
}
