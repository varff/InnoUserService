-- +goose Down
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_phone;