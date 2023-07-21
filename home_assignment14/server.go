package main

import (
	"github.com/gorilla/mux"
	"go-course/home_assignment14/controllers"
	"go-course/home_assignment14/middleware"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.IsValidApiKey)
	r.HandleFunc("/order", controllers.Create).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
