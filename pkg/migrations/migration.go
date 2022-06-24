package migration

import (

	"database/sql"


	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	dir = "pkg/migrations"
)


func MigrationUp(db *sql.DB) error {

	if err := goose.Up(db, dir); err != nil {
		return err
	}
	return nil
}
