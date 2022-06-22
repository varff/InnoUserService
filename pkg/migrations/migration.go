package migration

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"

	"InnoUserService/pkg/settings"
)

const (
	dir = "configs"
)

func UserMigrationUp(s *settings.DBSetting) error {
	connStr, err := settings.DBConnection(s)
	if err != nil {
		return err
	}
	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.Up(db, dir); err != nil {
		return err
	}
	return nil
}

func UserMigrationDown(s *settings.DBSetting) error {
	connStr, err := settings.DBConnection(s)
	if err != nil {
		return err
	}
	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.Down(db, dir); err != nil {
		return err
	}
	return nil
}
