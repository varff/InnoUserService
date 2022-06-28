package services

import (
	"errors"

	"InnoUserService/pkg/models"
	"InnoUserService/pkg/repo"
	"InnoUserService/pkg/settings"
)

type Service struct {
	repo     repo.IRepository
	settings *settings.AppSettings
}

type IService interface {
	Delete(phone int32) error
	Update(input models.RegisterModel) error
}

func (s Service) Delete(phone int32) error {
	if err := s.repo.DeleteUser(phone); err != nil {
		return errors.New("invalid phone number: " + err.Error())
	}
	return nil
}

func (s Service) Update(input models.RegisterModel) error {
	_, err := s.repo.GetUserByPhone(input.Phone)
	if err != nil {
		return errors.New("invalid phone number: " + err.Error())
	}
	pass, err := hashPassword(input.Pass)
	if err != nil {
		return errors.New("wrong password: " + err.Error())
	}
	if err = s.repo.UpdateUser(input.Name, pass, input.Email, input.Phone); err != nil {
		return errors.New("update error: " + err.Error())
	}
	return nil
}
