package dto

import "time"

// TodoCreateDTO info
// @Description Todo information create dto
type TodoCreateDTO struct {
	Task       string    `json:"task"`
	Status     string    `json:"status"`
	Project    string    `json:"project,omitempty"`
	StartDt    time.Time `json:"start_dt"`
	EndDt      time.Time `json:"end_dt"`
	User       string    `json:"user"`
	Department string    `json:"department,omitempty"`
} //@name TodoCreateDTO
