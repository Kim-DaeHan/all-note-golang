package dto

// UserCreateDTO info
// @Description User information create dto
type UserCreateDTO struct {
	GoogleID string `json:"google_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Verified *bool  `json:"verified,omitempty"`
	Provider string `json:"provider"`
	Photo    string `json:"photo"`
} //@name UserCreateDTO
