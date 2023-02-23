package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	serveHomePage(w, r)
}

var tplContainer *template.Template

var tplFunctions = template.FuncMap{
	"monthYYYY": monthYYYY,
}

func init() {
	tplContainer = template.Must(template.New("newTpl").Funcs(tplFunctions).ParseGlob("template/*"))
}

type user struct {
	Name       string
	Username   string
	password   string
	SignupDate time.Time
}

func main() {
	var newHandler myHandler
	http.ListenAndServe(":8080", newHandler)

}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
	list := map[string]string{
		"one": "hello",
		"two": "world",
	}

	Barin := user{
		Name:       "Barin Debnath",
		Username:   "barindebnath",
		password:   "password",
		SignupDate: time.Now(),
	}
	Zoro := user{
		Name:       "Roronoa Zoro",
		Username:   "roronoazoro",
		password:   "password",
		SignupDate: time.Now(),
	}
	Nico := user{
		"Nico Robin",
		"nicorobin",
		"password",
		time.Now(),
	}

	users := []user{Barin, Zoro, Nico}

	homeData := struct {
		Users []user
		List  map[string]string
	}{
		users,
		list,
	}

	err := tplContainer.ExecuteTemplate(w, "home.html", homeData)
	if err != nil {
		log.Fatalln(err)
	}

}

func monthYYYY(t time.Time) string {
	return t.Format("January 2006")
}
