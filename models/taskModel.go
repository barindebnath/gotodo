package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"task_id" bson:"_id"`
	Label     string             `json:"label" bson:"label"`
	IsDone    bool               `json:"is_done" bson:"is_done"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
