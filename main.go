package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	blockchain2 "go-course/home_assignmwnt13blockchain"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/transaction", addBlock).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

type RequestBody struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
}

type Response struct {
	Hash string `json:"hash"`
	Pow  int    `json:"pow"`
}

func addBlock(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to get request bode", http.StatusInternalServerError)
		return
	}

	blockchain := blockchain2.CreateBlockchain(2)

	newBlock := blockchain.AddBlock(
		reqBody.From,
		reqBody.To,
		reqBody.Amount,
	)

	response := Response{
		Hash: newBlock.GetHash(),
		Pow:  newBlock.GetPow(),
	}
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
