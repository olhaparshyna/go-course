package requests

import (
	"encoding/json"
	"net/http"
)

type CreateRequestData struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Items []string `json:"items"`
}

func ValidateCreateRequestData(w http.ResponseWriter, r *http.Request) *CreateRequestData {
	var requestData CreateRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
	}

	if requestData.Name == "" || requestData.Email == "" || len(requestData.Items) == 0 {

		http.Error(w, "Missing required fields", http.StatusBadRequest)
	}

	return &requestData
}
