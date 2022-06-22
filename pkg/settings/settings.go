package settings

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type DBSetting struct {
	DBPort     string
	DBHost     string
	DBName     string
	DBUser     string
	DBPassword string
	SSLMode    string
}

func NewDBSetting() (*DBSetting, error) {
	viper.SetConfigType("env")
	viper.SetConfigName("conf")
	viper.AddConfigPath("configs/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &DBSetting{}
	s.DBUser = viper.GetString("USERID")
	s.DBPassword = viper.GetString("USERPASS")
	s.DBPort = viper.GetString("USERPORT")
	s.DBName = viper.GetString("USERDB")
	s.DBHost = viper.GetString("USERHOSTNAME")
	s.SSLMode = viper.GetString("USERSSL")
	return s, nil
}

func DBConnection(setting *DBSetting) (string, error) {
	UserID, err := GetEnvDefault(setting.DBUser, "user")
	if err != nil {
		return "", err
	}
	Pass, err := GetEnvDefault(setting.DBPassword, "secret")
	if err != nil {
		return "", err
	}
	Port, err := GetEnvDefault(setting.DBPort, "5432")
	if err != nil {
		return "", err
	}
	Db, err := GetEnvDefault(setting.DBName, "postgres")
	if err != nil {
		return "", err
	}
	Host, err := GetEnvDefault(setting.DBHost, "localhost")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("user=" + UserID + " password=" + Pass + " host=" + Host + " port=" + Port + " database=" + Db), nil
}

func GetEnvDefault(key, defaultValue string) (string, error) {
	if key == "" {
		if defaultValue == "" {
			return "", errors.New("environment variable isn't set")
		}
		return defaultValue, nil
	}

	return key, nil
}
