-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE transaction_method AS ENUM ('cash', 'transfer', 'ewallet');
CREATE TYPE transaction_status AS ENUM ('pending', 'failed', 'paid');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    rental_id INT NOT NULL,
    amount INT NOT NULL,
    method transaction_method NOT NULL,
    status transaction_status DEFAULT 'pending',
    proof TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS transaction_method;
DROP TYPE IF EXISTS transaction_status;

-- +migrate StatementEnd
