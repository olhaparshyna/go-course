package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-course/home_assignment16/database"
	"go-course/home_assignment16/handlers"
	"go-course/home_assignment16/middleware"
	"go-course/home_assignment16/requests"
	"log"
	"net/http"
)

func main() {
	db := database.InitDB()
	conf, storage, err := initApp(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := mux.NewRouter()
	r.Use(middleware.IsValidApiKey(conf))
	r.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		responseData, err := handlers.ProductList(storage.ProductRepo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(responseData)
	}).Methods(http.MethodGet)

	r.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		requestData := requests.ValidateCreateRequestData(w, r)

		responseData, err := handlers.Create(*storage, *requestData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(responseData)
	}).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}
