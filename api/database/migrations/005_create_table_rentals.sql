-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE rental_status AS ENUM ('pending', 'active', 'finished', 'cancelled');

CREATE TABLE rentals (
    id SERIAL PRIMARY KEY,
    room_id INT NOT NULL,
    tenant_id INT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    duration_months INT NOT NULL,
    status rental_status DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS rentals;
DROP TYPE IF EXISTS rental_status;

-- +migrate StatementEnd
