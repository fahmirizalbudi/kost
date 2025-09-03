-- +migrate Up
-- +migrate StatementBegin

CREATE TYPE transaction_method AS ENUM ('cash', 'transfer', 'ewallet');
CREATE TYPE transaction_purpose AS ENUM ('new', 'renewal');

-- TODO: add dormitory_price column, add purpose column ('new', 'renewal')

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    rental_id INT NOT NULL,
    dormitory_price INT NOT NULL,
    month_paid INT NOT NULL,
    amount INT NOT NULL,
    method transaction_method NOT NULL,
    purpose transaction_purpose DEFAULT 'new',
    proof TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS transaction_method;
DROP TYPE IF EXISTS transaction_purpose;

-- +migrate StatementEnd
