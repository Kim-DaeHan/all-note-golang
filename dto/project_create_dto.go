package dto

import "time"

// ProjectCreateDTO info
// @Description Project information create dto
type ProjectCreateDTO struct {
	Name    string    `json:"name"`
	StartDt time.Time `json:"start_dt"`
	EndDt   time.Time `json:"end_dt,omitempty"`
} //@name ProjectCreateDTO
