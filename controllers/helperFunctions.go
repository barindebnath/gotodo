package controllers

import (
	"context"
	"encoding/base64"
	"errors"
	"log"
	"net/http"

	"github.com/barindebnath/gotodo/database"
	"github.com/barindebnath/gotodo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getIdFromCookie(r *http.Request) (primitive.ObjectID, string) {
	var id primitive.ObjectID
	var error string

	cookie, err := r.Cookie("userSession")

	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			error = "cookie not found"
		default:
			log.Println(err)
			error = "Server error: Error finding cookie"
		}
	} else {
		value, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			log.Println(err)
			error = "Server error: Error decoding cookie"
		} else {
			id, err = primitive.ObjectIDFromHex(string(value))
			if err != nil {
				log.Println(err)
				error = "Server error: Error converting cookie value to primitive.ObjectID"
			}
		}
	}

	return id, error
}

func getUserById(id primitive.ObjectID) (models.User, error) {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := models.User{}

	collection := database.OpenCollection(database.Client)
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	return result, err
}

func getUserFromCookie(r *http.Request) (models.User, string) {
	id, error := getIdFromCookie(r)

	user, err := getUserById(id)
	if err != nil {
		error = "User not found"
	}

	return user, error
}

func removeUserSession(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("userSession")
	if err == nil {
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
}
