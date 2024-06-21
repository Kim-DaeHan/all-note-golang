package services

import (
	"github.com/Kim-DaeHan/all-note-golang/dto"
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type TodoService interface {
	GetAllTodo() ([]models.Todo, error)
	GetTodo(id string) (*models.Todo, error)
	GetTodoByUser(userId string) ([]models.Todo, error)
	CreateTodo(dto *dto.TodoCreateDTO) error
	UpdateTodo(id string, dto *dto.TodoUpdateDTO) (*models.Todo, error)
	DeleteTodo(id string) error
}
