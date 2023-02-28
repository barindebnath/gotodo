package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/barindebnath/gotodo/router"
)

func main() {
	mux := router.Router()
	log.Fatal(http.ListenAndServe(":8080", mux))
	fmt.Println("Listning at port 8080 . . .")
}
