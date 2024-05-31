package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Email       string             `json:"email,omitempty" validate:"required"`
	Role        string             `json:"role,omitempty" validate:"required"`
	Date        primitive.DateTime `json:"date,omitempty"`
	JoiningDate string             `json:"joining_date,omitempty" validate:"required"`
}
