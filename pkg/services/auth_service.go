package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"InnoUserService/pkg/apperrors"
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
		return "", apperrors.Wrapper(apperrors.ErrWrongPass, err)
	}
	token, err := s.CreateToken(user.Phone, s.settings.TTLMinutes)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) Register(input models.RegisterModel) (string, error) {
	_, err := s.repo.GetUserByPhone(input.Phone)
	if err == nil {
		return "", apperrors.Wrapper(apperrors.ErrPhoneTaken, err)
	}
	pass, err := hashPassword(input.Pass)
	if err != nil {
		return "", apperrors.Wrapper(apperrors.ErrWrongPass, err)
	}
	err = s.repo.AddUser(input.Name, pass, input.Email, input.Phone)
	token, err := s.CreateToken(input.Phone, s.settings.TTLMinutes)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Service) CreateToken(userPhone int32, ttl time.Duration) (string, error) {
	var err error
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserPhone: userPhone,
	})
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", apperrors.Wrapper(apperrors.ErrTokenSigned, err)
	}
	return token, nil
}

func (s Service) ParseToken(accessToken string) (int32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrors.Wrapper(apperrors.ErrSingingMethod, nil)
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return 0, apperrors.Wrapper(apperrors.ErrTokenParsing, err)
	}
	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return 0, apperrors.Wrapper(apperrors.ErrWrongClaims, nil)
	}
	return claims.UserPhone, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
