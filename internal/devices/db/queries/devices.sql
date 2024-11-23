-- name: CreateDevice :one
INSERT INTO devices (name, type, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetDeviceByID :one
SELECT id, name, type, status, created_at, updated_at
FROM devices
WHERE id = $1;

-- name: UpdateDeviceStatus :exec
UPDATE devices
SET status = $2, updated_at = $3
WHERE id = $1;

-- name: ListDevices :many
SELECT id, name, type, status, created_at, updated_at
FROM devices;

-- name: ListDevicesByStatus :many
SELECT id, name, type, status, created_at, updated_at
FROM devices
WHERE status = $1;
