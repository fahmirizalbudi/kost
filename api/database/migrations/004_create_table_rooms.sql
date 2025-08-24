-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE room_status AS ENUM ('available', 'rented');

CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    boarding_house_id INT,
    room_number VARCHAR(50),
    status room_status DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS rooms;
DROP TYPE IF EXISTS room_status;

-- +migrate StatementEnd
