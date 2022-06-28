package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"InnoUserService/pkg/models"
)

type IAuth interface {
	Login(input models.LoginModel) (string, error)
	Register(input models.RegisterModel) (string, error)
	CreateToken(userid uint64, TTL time.Duration) (string, error)
}

func (s Service) Login(input models.LoginModel) (string, error) {
	var user models.User
	user, err := s.repo.GetUserByPhone(input.Phone)
	if err != nil {
		return "", err
	}
	if !checkPasswordHash(input.Pass, user.Password) {
		return "", errors.New("wrong password: " + err.Error())
	}
	token, err := CreateToken(uint64(user.Phone), s.settings.TTLMinutes)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) Register(input models.RegisterModel) (string, error) {
	var user models.User
	user, err := s.repo.GetUserByPhone(input.Phone)
	if err != nil {
		return "", errors.New("invalid phone number: " + err.Error())
	}
	pass, err := hashPassword(input.Pass)
	if err != nil {
		return "", errors.New("wrong password: " + err.Error())
	}
	err = s.repo.AddUser(input.Name, pass, input.Email, input.Phone)
	token, err := CreateToken(uint64(user.Phone), s.settings.TTLMinutes)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateToken(userid uint64, TTL time.Duration) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(TTL).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
