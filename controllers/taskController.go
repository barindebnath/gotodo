package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/barindebnath/gotodo/database"
	"github.com/barindebnath/gotodo/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	user, errorString := getUserFromCookie(r)

	if errorString != "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": errorString,
		})
	} else {
		json.NewEncoder(w).Encode(user.Tasks)
	}
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	id, errorString := getIdFromCookie(r)

	if errorString != "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": errorString,
		})
	} else {
		r.ParseMultipartForm(1)
		var description string

		for key, value := range r.MultipartForm.Value {
			if key == "description" {
				description = value[0]
			}
		}

		if len(description) == 0 {
			errorString = "Enter a description."
			json.NewEncoder(w).Encode(map[string]string{
				"error": errorString,
			})
		} else {
			var newTask = models.Task{
				ID:        uuid.NewString(),
				Label:     description,
				IsDone:    false,
				CreatedAt: time.Now(),
			}

			// user.Tasks = append(user.Tasks, newTask)

			filter := bson.M{"_id": bson.M{"$eq": id}}
			update := bson.M{"$push": bson.M{"tasks": newTask}}

			collection := database.OpenCollection(database.Client)
			_, err := collection.UpdateOne(context.Background(), filter, update)
			if err != nil {
				log.Println(err)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Unable to add task.",
				})
			} else {
				json.NewEncoder(w).Encode(newTask)
			}
		}
	}
}

func MarkAsDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	id, errorString := getIdFromCookie(r)

	if errorString != "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": errorString,
		})
	} else {
		var taskId string
		var isDone string

		// extract task id from url
		vars := mux.Vars(r)
		taskId = vars["id"]
		if len(taskId) < 1 {
			log.Println("Task ID is missing")
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Task ID is missing",
			})
			return
		}

		// extract is_done from request body
		r.ParseMultipartForm(1)
		for key, value := range r.MultipartForm.Value {
			if key == "is_done" {
				isDone = value[0]
			}
		}

		if len(isDone) == 0 {
			json.NewEncoder(w).Encode(map[string]string{
				"error": "'is_done' is missing.",
			})
		} else {
			filter := bson.M{"_id": id, "tasks.task_id": taskId}
			update := bson.M{"$set": bson.M{"tasks.$.is_done": (isDone == "1")}}

			collection := database.OpenCollection(database.Client)
			_, err := collection.UpdateOne(context.Background(), filter, update)

			if err != nil {
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Unable to update task.",
				})
				log.Println(err)
				return
			}

			json.NewEncoder(w).Encode(map[string]string{
				"success": "Task updated successfully.",
			})

		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	id, errorString := getIdFromCookie(r)

	if errorString != "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": errorString,
		})
	} else {
		var taskId string

		// extract task id from url
		vars := mux.Vars(r)
		taskId = vars["id"]
		if len(taskId) < 1 {
			log.Println("Task ID is missing")
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Task ID is missing",
			})
			return
		}
		filter := bson.M{"_id": id}
		change := bson.M{"$pull": bson.M{"tasks": bson.M{"task_id": taskId}}}

		collection := database.OpenCollection(database.Client)
		_, err := collection.UpdateOne(context.Background(), filter, change)
		if err != nil {
			log.Println(err)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Unable to delete task.",
			})
		} else {
			json.NewEncoder(w).Encode(map[string]string{
				"success": "Task deleted successfully.",
			})
		}
	}
}
