package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUser(id string) (*models.User, error)
	CreateUser(dto *dto.UserCreateDTO) error
	UpsertUser(dto *dto.UserUpdateDTO) (*models.User, error)
}
