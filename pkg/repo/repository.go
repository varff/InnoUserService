package repo

import (
	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"

	"InnoUserService/pkg/models"
	"InnoUserService/pkg/settings"
)

type IRepository interface {
	GetUserByPhone(int32) (models.User, error)
	GetUserRateByPhone(int32) (models.User, error)
	AddUser(Name, Password, Email string, Phone int32) (bool, error)
}

type Repository struct {
	*sqlx.DB
}

func NewRepository(dbSettings *settings.DBSetting) (*Repository, error) {
	connStr, err := settings.DBConnection(dbSettings)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repository{DB: db}, nil
}

func (db Repository) GetUserByPhone(phone int32) (models.User, error) {
	var result models.User
	que := "SELECT * FROM users where phone = $1;"
	row := db.QueryRow(que, phone)
	err := row.Scan(&result)
	if err == pgx.ErrNoRows {
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (db Repository) GetUserRateByPhone(phone int32) (int32, error) {
	var result int32
	que := "SELECT rate FROM users where phone = $1;"
	row := db.QueryRow(que, phone)
	err := row.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (db Repository) AddUser(Name, Password, Email string, Phone int32) (bool, error) {
	que := "insert into users (\"name\", pass, phone, email, rate, analyst) values($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(que, Name, Password, Phone, Email, 0, false)
	if err != nil {
		return false, err
	}
	return true, nil
}
