package dto

import "time"

// ProjectUpdateDTO info
// @Description Project information update dto
type ProjectUpdateDTO struct {
	Name    string    `json:"name,omitempty"`
	StartDt time.Time `json:"start_dt,omitempty"`
	EndDt   time.Time `json:"end_dt,omitempty"`
} //@name ProjectUpdateDTO
