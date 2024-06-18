package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project info
// @Description Project information
type Project struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	StartDt   time.Time          `bson:"start_dt" json:"start_dt"`
	EndDt     time.Time          `bson:"end_dt" json:"end_dt"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
} //@name Project
