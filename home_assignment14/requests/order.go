package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequestData struct {
	UserId int      `json:"userId"`
	Items  []string `json:"items"`
}

func ValidateCreateRequestData(w http.ResponseWriter, r *http.Request) *CreateRequestData {
	var requestData CreateRequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
	}

	if requestData.UserId == 0 || len(requestData.Items) == 0 {
		fmt.Println(requestData.UserId)
		fmt.Println(requestData.Items)

		http.Error(w, "Missing required fields", http.StatusBadRequest)
	}

	return &requestData
}
