package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Department info
// @Description Department information
type Department struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	ParantId primitive.ObjectID `bson:"parant_id" json:"parant_id"`
} //@name Department
