-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    Id      serial primary key,
    Name    varchar(32),
    Pass    varchar(256),
    Phone   int unique,
    Email   varchar(32),
    Rate    smallint,
    Analyst bool
);

CREATE INDEX IF NOT EXISTS idx_phone ON users(Phone);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_phone;