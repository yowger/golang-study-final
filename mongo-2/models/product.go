package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	// ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// Name        string             `bson:"name" json:"name" validate:"required"`
	// Description string             `bson:"description" json:"description"`
	// Price       float64            `bson:"price" json:"price" validate:"required,gt=0"`
	ID          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description"`
	Price       float64            `json:"price" validate:"required,gt=0"`
}
