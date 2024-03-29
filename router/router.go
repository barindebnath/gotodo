package router

import (
	controller "github.com/barindebnath/gotodo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// pages
	router.HandleFunc("/", controller.HomePage).Methods("GET")
	router.HandleFunc("/employees", controller.EmployeesPage).Methods("GET")
	router.HandleFunc("/signup", controller.SignupPage).Methods("GET")
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	router.HandleFunc("/logout", controller.LogoutPage).Methods("GET")

	// user apis routes
	router.HandleFunc("/api/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/api/login", controller.Login).Methods("POST")

	// task apis route
	router.HandleFunc("/api/tasks", controller.GetTasks).Methods("GET")
	router.HandleFunc("/api/task", controller.AddTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", controller.MarkAsDone).Methods("PUT")
	router.HandleFunc("/api/task/{id}", controller.DeleteTask).Methods("DELETE")

	// static files
	router.HandleFunc("/favicon-16x16.ico", controller.FaviconHandler)

	return router
}
