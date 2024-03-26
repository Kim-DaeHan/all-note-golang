package dto

type UserUpdateDTO struct {
	GoogleID string `json:"google_id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Verified *bool  `json:"verified,omitempty"`
	Provider string `json:"provider"`
	Photo    string `json:"photo"`
}
