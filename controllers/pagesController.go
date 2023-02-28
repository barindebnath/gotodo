package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/barindebnath/gotodo/database"
	"github.com/barindebnath/gotodo/models"
	"github.com/barindebnath/gotodo/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	err := template.Templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	err := template.Templates.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func EmplyeesPage(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	users := getUsers()

	err := template.Templates.ExecuteTemplate(w, "users.html", users)
	if err != nil {
		log.Fatalln(err)
	}
	// json.NewEncoder(w).Encode(users)
}

// get all users from db
func getUsers() []models.User {
	var collection *mongo.Collection = database.OpenCollection(database.Client)
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer cursor.Close(context.Background())
	return users
}
