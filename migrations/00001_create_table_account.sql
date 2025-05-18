-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    name VARCHAR(255) NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    currency CHAR(3) NOT NULL CHECK (currency ~ '^[A-Z]{3}$')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS account;
-- +goose StatementEnd
