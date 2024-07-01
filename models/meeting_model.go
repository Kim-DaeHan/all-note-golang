package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Meeting info
// @Description Meeting information
type Meeting struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Title        string             `bson:"title" json:"title"`
	Description  string             `bson:"description" json:"description"`
	Participants []Participant      `bson:"participants" json:"participants"`
	StartDt      time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt        time.Time          `bson:"end_dt" json:"end_dt"`
	Location     string             `bson:"location" json:"location"`
	User         primitive.ObjectID `bson:"created_by" json:"created_by"`
	UserInfo     []User             `bson:"created_by_info,omitempty" json:"created_by_info,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Meeting

type Participant struct {
	User     primitive.ObjectID `bson:"participant" json:"participant"`
	UserInfo []User             `bson:"participant_info,omitempty" json:"participant_info,omitempty"`
}
