package models

type User struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Title    string `json:"title" validate:"required"`
}
