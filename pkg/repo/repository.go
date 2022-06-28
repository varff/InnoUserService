package repo

import (
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jmoiron/sqlx"

	migration "InnoUserService/pkg/migrations"
	"InnoUserService/pkg/models"
	"InnoUserService/pkg/settings"
)

type IRepository interface {
	AddUser(name, password, email string, phone int32) error
	DeleteUser(phone int32) error
	GetUserByMail(mail string) (models.User, error)
	GetUserByPhone(phone int32) (models.User, error)
	GetUserId(phone int32) (int32, error)
	UpdateUser(name, password, email string, phone int32) error
}

type Repository struct {
	*sqlx.DB
}

func NewRepository(dbSetting *settings.DBSetting) (*Repository, error) {
	connStr := fmt.Sprintf("user=" + dbSetting.DBUser + " password=" + dbSetting.DBPassword + " host=" + dbSetting.DBHost + " port=" + dbSetting.DBPort + " database=" + dbSetting.DBName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = migration.MigrationUp(db.DB)
	if err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

func (db Repository) AddUser(name, password, email string, phone int32) error {
	query := "insert into users (\"name\", pass, phone, email, analyst) values($1, $2, $3, $4, $5)"
	_, err := db.Exec(query, name, password, phone, email, false)
	if err != nil {
		return err
	}
	return nil
}

func (db Repository) DeleteUser(phone int32) error {
	query := "delete from users where phone=$1"
	_, err := db.Exec(query, phone)
	if err != nil {
		return err
	}
	return nil
}

func (db Repository) GetUserByMail(mail string) (models.User, error) {
	var result models.User
	query := "select * from users where email = $1;"
	row := db.QueryRow(query, mail)
	err := row.Scan(&result)
	if err == pgx.ErrNoRows {
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (db Repository) GetUserByPhone(phone int32) (models.User, error) {
	var result models.User
	query := "select * from users where phone = $1;"
	row := db.QueryRow(query, phone)
	err := row.Scan(&result)
	if err == pgx.ErrNoRows {
		return result, err
	}
	if err != nil {
		return result, err
	}
	return result, nil
}

func (db Repository) GetUserId(phone int32) (int32, error) {
	var result int32
	query := "select id from users where phone = $1;"
	row := db.QueryRow(query, phone)
	err := row.Scan(&result)
	return result, err
}

func (db Repository) UpdateUser(name, password, email string, phone int32) error {
	query := "update users set name=$1, pass=$2, email=$3 where phone=$4"
	_, err := db.Exec(query, name, password, email, phone)
	return err
}
