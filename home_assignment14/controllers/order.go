package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-course/home_assignment14/models"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var orderData models.Order
	err := json.NewDecoder(r.Body).Decode(&orderData)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if orderData.UserId == 0 || len(orderData.Items) == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	orderData.Id = uuid.New()

	processor := models.NewProcessor(&models.Storage)

	err = processor.ProcessOrder(orderData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orderData)
}
