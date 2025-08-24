-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE boarding_houses (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    description TEXT,
    price INT NOT NULL,
    facilities TEXT,
    google_maps TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS boarding_houses;

-- +migrate StatementEnd
