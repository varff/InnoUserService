package migration

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"

	"InnoUserService/pkg/settings"
)

func UserMigrationUp(s *settings.DBSetting) error {
	connStr, err := settings.UserConString(s)
	if err != nil {
		return err
	}
	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	dir := "pkg/migrations/"

	if err := goose.Up(db, dir); err != nil {
		return err
	}
	return nil
}

func UserMigrationDown(s *settings.DBSetting) error {
	connStr, err := settings.UserConString(s)
	if err != nil {
		return err
	}
	db, err := goose.OpenDBWithDriver("pgx", connStr)
	if err != nil {
		return err
	}
	defer db.Close()
	dir := "pkg/migrations/"

	if err := goose.Down(db, dir); err != nil {
		return err
	}
	return nil
}
