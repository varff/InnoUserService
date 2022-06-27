package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"InnoUserService/pkg/repo"
	"InnoUserService/pkg/settings"
)

type Service struct {
	repo     repo.IRepository
	settings *settings.AppSettings
}

type IService interface {
	Delete(c *gin.Context) error
}

func (s Service) Delete(c *gin.Context) error {
	var phone int32
	if err := c.ShouldBindJSON(&phone); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return err
	}
	err := s.repo.DeleteUser(phone)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid phone number")
		return err
	}
	c.JSON(http.StatusOK, "Success")
	return err
}
