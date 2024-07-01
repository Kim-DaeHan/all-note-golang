package services

import (
	"github.com/Kim-DaeHan/all-note-golang/models"
)

type MeetingService interface {
	// GetAllTodo() ([]models.Todo, error)
	GetMeeting(id string) (*models.Meeting, error)
	// GetTodoByUser(userId string) ([]models.Todo, error)
	// CreateTodo(dto *dto.TodoCreateDTO) error
	// UpdateTodo(id string, dto *dto.TodoUpdateDTO) (*models.Todo, error)
	// DeleteTodo(id string) error
}
