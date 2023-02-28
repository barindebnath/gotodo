package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type task = struct {
	Id      string `json:"id"`
	Detail  string `json:"detail"`
	Checked bool   `json:"checked"`
}

var tasks = []task{}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	json.NewEncoder(w).Encode(tasks)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	r.ParseMultipartForm(1)
	var description string

	for key, value := range r.MultipartForm.Value {
		if key == "description" {
			description = value[0]
		}
	}

	if len(description) > 0 {
		var newTask = task{
			uuid.NewString(),
			description,
			false,
		}
		tasks = append(tasks, newTask)

		json.NewEncoder(w).Encode(newTask)
	}
}

func MarkAsDone(w http.ResponseWriter, r *http.Request) {

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

// func checkNilErr(err error) {
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		panic(err)
// 	}
// }
