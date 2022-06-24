-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    Id      serial primary key,
    Name    varchar(32),
    Pass    varchar(256),
    Phone   int unique,
    Email   varchar(32),
    Analyst bool
);

CREATE INDEX IF NOT EXISTS idx_phone ON users(Phone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_phone;
-- +goose StatementEnd
