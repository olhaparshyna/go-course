package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-course/home_assignment14/handlers"
	"go-course/home_assignment14/middleware"
	"go-course/home_assignment14/requests"
	"log"
	"net/http"
)

func main() {
	conf, processor, err := initApp()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.Use(middleware.IsValidApiKey(conf))
	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		requestData := requests.ValidateCreateRequestData(w, r)

		responseData, err := handlers.Create(*processor, *requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(responseData)
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
