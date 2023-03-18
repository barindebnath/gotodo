package models

import (
	"time"
)

type Task struct {
	ID        string    `json:"task_id" bson:"task_id"`
	Label     string    `json:"label" bson:"label"`
	IsDone    bool      `json:"is_done" bson:"is_done"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
