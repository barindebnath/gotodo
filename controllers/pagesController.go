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

type HomePageData struct {
	User models.User
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	user, errorString := getUserFromCookie(r)

	if errorString != "" {
		err := template.Templates.ExecuteTemplate(w, "home.html", nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		var data = HomePageData{
			user,
		}
		err := template.Templates.ExecuteTemplate(w, "home.html", data)
		if err != nil {
			log.Println(err)
		}
	}

}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	err := template.Templates.ExecuteTemplate(w, "signup.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	err := template.Templates.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Println(err)
	}
}

func LogoutPage(w http.ResponseWriter, r *http.Request) {
	removeUserSession(w, r) // can implement go routine here
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type EmployeesPageData struct {
	Users            []models.User
	CurrentUserEmail string
}

func EmployeesPage(w http.ResponseWriter, r *http.Request) {
	user, errorString := getUserFromCookie(r)

	if errorString != "" {
		err := template.Templates.ExecuteTemplate(w, "users.html", nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		var data = EmployeesPageData{
			getUsers(),
			user.Email,
		}

		err := template.Templates.ExecuteTemplate(w, "users.html", data)
		if err != nil {
			log.Println(err)
		}
	}
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

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/images/favicon-16x16.ico")
}
