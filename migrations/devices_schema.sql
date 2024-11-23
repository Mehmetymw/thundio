CREATE TYPE device_status AS ENUM ('Active','Inactive','Maintenance','Faulty','Disconnected');

CREATE TABLE devices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    status device_status NOT NULL DEFAULT 'Inactive',
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);