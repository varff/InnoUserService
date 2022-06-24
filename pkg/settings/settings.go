package settings

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

)

type DBSetting struct {
	DBPort     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	SSLMode    string
}


type AppSettings struct {
	Salt       string
	Port       int32
	TTLMinutes time.Duration
}

func NewAppSettings() (*AppSettings, error) {
	s := &AppSettings{}
	var err error
	s.Salt, err = GetEnvDefault("JWT_SALT", "safeless")

	if err != nil {
		return s, err
	}
	portStr, err := GetEnvDefault("APP_PORT", "8000")
	if err != nil {
		return s, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 8000
	}
	s.Port = int32(port)
	TTLStr, err := GetEnvDefault("APP_TTL", "8000")
	if err != nil {
		return s, err
	}
	TTL, err := strconv.Atoi(TTLStr)
	if err != nil {
		TTL = 60
	}
	s.TTLMinutes, err = time.ParseDuration(fmt.Sprintf("%dm", TTL))
	if err != nil {
		s.TTLMinutes = time.Hour
	}
	return &AppSettings{}, nil
}


func NewDBSetting() (*DBSetting, error) {
	s := &DBSetting{}
	var err error
	s.DBUser, err = GetEnvDefault("USERID", "user")

	if err != nil {
		return s, err
	}
	s.DBPassword, err = GetEnvDefault("USERPASS", "secret")
	if err != nil {
		return s, err
	}
	s.DBPort, err = GetEnvDefault("USERPORT", "5432")
	if err != nil {
		return s, err
	}
	s.DBName, err = GetEnvDefault("USERDB", "postgres")
	if err != nil {
		return s, err
	}
	s.DBHost, err = GetEnvDefault("USERHOSTNAME", "localhost")
	if err != nil {
		return s, err
	}
	s.SSLMode, err = GetEnvDefault("USERSSL", "false")
	if err != nil {
		return s, err
	}
	return s, nil
}

func GetEnvDefault(key, defaultValue string) (string, error) {

	value := os.Getenv(key)

	if key == "" {
		if defaultValue == "" {
			return "", errors.New("environment variable isn't set")
		}
		return defaultValue, nil
	}


	return value, nil

}
