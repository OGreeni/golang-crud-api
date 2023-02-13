package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name,omitempty" bson:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name,omitempty" bson:"last_name,omitempty" validate:"required"`
	Gender     string             `json:"gender,omitempty" bson:"gender,omitempty" validate:"required"`
	BirthDate string             `json:"birth_date,omitempty" bson:"birth_date,omitempty" validate:"required"`
}
