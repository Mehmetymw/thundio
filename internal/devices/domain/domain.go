package domain

import (
	"database/sql"
	"time"
)

type DeviceID int32

type Device struct {
	ID        DeviceID     `json:"id"`
	UniqueID  string       `json:"unique_id"`
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Status    DeviceStatus `json:"Status"`
	LastSeen  sql.NullTime `json:"last_seen"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (d *Device) Activate() {
	d.Status = Active
}

func (d *Device) Deactivate() {
	d.Status = Inactive
}
