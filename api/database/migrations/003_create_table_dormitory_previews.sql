-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE dormitory_previews (
    id SERIAL PRIMARY KEY,
    dormitory_id INT,
    url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS boarding_house_details;

-- +migrate StatementEnd
