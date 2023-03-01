package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/barindebnath/gotodo/database"
	"github.com/barindebnath/gotodo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var isValid = true
	var errorMsg = map[string]string{
		"first_name": "",
		"last_name":  "",
		"email":      "",
		"password":   "",
	}

	r.ParseMultipartForm(1)
	var first_name string
	var last_name string
	var email string
	var password string

	for key, value := range r.MultipartForm.Value {
		switch key {
		case "first_name":
			if validateString(value[0], &isValid) {
				first_name = value[0]
			} else {
				errorMsg["first_name"] = "Enter your first name."
			}

		case "last_name":
			if validateString(value[0], &isValid) {
				last_name = value[0]
			} else {
				errorMsg["last_name"] = "Enter your last name."
			}

		case "email":
			if validateEmail(value[0], &isValid) {
				email = value[0]
			} else {
				errorMsg["email"] = "Enter correct email id."
			}

		case "password":
			if validateString(value[0], &isValid) {
				password = value[0]
			} else {
				errorMsg["password"] = "Enter a new password."
			}

		default:
			fmt.Println("KEY NOT FOUND")
		}
	}

	if isValid {
		var newUser = models.User{
			ID:        primitive.NewObjectID(),
			FirstName: first_name,
			LastName:  last_name,
			Email:     email,
			Password:  password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		collection := database.OpenCollection(database.Client)
		_, err := collection.InsertOne(context.Background(), newUser)

		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User created successfully",
		})
	} else {
		json.NewEncoder(w).Encode(errorMsg)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func validateString(value string, isValid *bool) bool {
	if len(strings.Trim(value, " ")) > 0 {
		return true
	} else {
		*isValid = false
		return false
	}
}

func validateEmail(value string, isValid *bool) bool {
	_, err := mail.ParseAddress(value)
	// collection := database.OpenCollection(database.Client)
	// _, notFound := collection.Find(func()bson.D{{email:value} })

	if err == nil {
		return true
	} else {
		*isValid = false
		return false
	}
}
