package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID        primitive.ObjectID `json:"user_id" bson:"_id"`
	FirstName string             `json:"first_name" validate:"required,min=2,max=100" bson:"first_name"`
	LastName  string             `json:"last_name" validate:"required,min=2,max=100" bson:"last_name"`
	Password  string             `json:"password" validate:"required,min=6" bson:"password"`
	Email     string             `json:"email" validate:"email,required" bson:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
