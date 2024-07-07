package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Department info
// @Description Department information
type Department struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	ParentId primitive.ObjectID `bson:"parent_id,omitempty" json:"parent_id,omitempty"`
} //@name Department
