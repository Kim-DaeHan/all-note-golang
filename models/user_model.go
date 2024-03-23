package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	GoogleID  string             `bson:"google_id"`
	Email     string             `bson:"email"`
	UserName  string             `bson:"user_name"`
	Verified  *bool              `bson:"verified,omitempty"`
	Provider  string             `bson:"provider"`
	Photo     string             `bson:"photo"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
