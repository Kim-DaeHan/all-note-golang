package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User info
// @Description User information
type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	GoogleID  string             `bson:"google_id" json:"google_id"`
	Email     string             `bson:"email" json:"email"`
	UserName  string             `bson:"user_name" json:"user_name"`
	Position  string             `bson:"position,omitempty" json:"position,omitempty"`
	Verified  *bool              `bson:"verified,omitempty" json:"verified,omitempty"`
	Provider  string             `bson:"provider" json:"provider"`
	Photo     string             `bson:"photo" json:"photo"`
	Team      primitive.ObjectID `bson:"team,omitempty" json:"team,omitempty"`
	TeamInfo  []Team             `bson:"team_info,omitempty" json:"team_info,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
} //@name User
