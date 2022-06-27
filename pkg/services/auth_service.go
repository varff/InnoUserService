package services

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"InnoUserService/pkg/models"
)

type IAuth interface {
	Login(c *gin.Context) (string, error)
	Register(c *gin.Context) (string, error)
}

func (s Service) Login(c *gin.Context) (string, error) {
	var (
		user models.User
		u    models.LoginModel
	)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return "", err
	}
	user, err := s.repo.GetUserByPhone(u.Phone)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid phone number")
		return "", err
	}

	if checkPasswordHash(u.Pass, user.Password) {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return "", nil
	}
	token, err := createToken(uint64(user.Phone), s.settings.TTLMinutes)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return "", err
	}
	c.JSON(http.StatusOK, token)
	return token, nil
}

func (s Service) Register(c *gin.Context) (string, error) {
	var (
		user models.User
		u    models.RegisterModel
	)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return "", err
	}
	user, err := s.repo.GetUserByPhone(u.Phone)
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid phone number")
		return "", err
	}
	pass, err := hashPassword(u.Pass)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid pass")
	}
	err = s.repo.AddUser(u.Name, pass, u.Email, u.Phone)
	token, err := createToken(uint64(user.Phone), s.settings.TTLMinutes)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return "", err
	}
	c.JSON(http.StatusOK, token)
	return token, nil
}

func createToken(userid uint64, TTL time.Duration) (string, error) {
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
