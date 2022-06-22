package main

import (
	"log"

	migration "InnoUserService/pkg/migrations"
	"InnoUserService/pkg/settings"
)

func main() {
	dbSettings, err := settings.NewDBSetting()
	if err != nil {
		log.Fatalf("DB settings loading error: %s: ", err)
	}
	err = migration.UserMigrationUp(dbSettings)
	if err != nil {
		log.Fatalf("migration up error: %s", err)
	}

}
