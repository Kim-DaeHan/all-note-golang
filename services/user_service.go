package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	CreateUser(dto dto.UserCreateDTO) (*mongo.InsertOneResult, error)
	UpsertUser(email string, dto dto.UserUpdateDTO) (*models.User, error)
}
