package main

import (
	"log"

	"InnoUserService/pkg/repo"
	"InnoUserService/pkg/settings"
)

func main() {

	dbSettings, err := settings.NewDBSetting()

	if err != nil {
		log.Fatalf("DB settings loading error: %s: ", err)
	}
	_, err = repo.NewRepository(dbSettings)
	if err != nil {
		log.Fatalf("DB connection loading error: %s: ", err)
	}

}
