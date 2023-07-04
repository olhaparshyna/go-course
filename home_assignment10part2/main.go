package main

import (
	"encoding/json"
	google_translate "github.com/gilang-as/google-translate"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/translate", translate).
		Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func translate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type RequestBody struct {
		Lang string `json:"lang"`
		Text string `json:"text"`
	}

	var reqBody RequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Failed to get request body", http.StatusInternalServerError)
		return
	}

	lang := reqBody.Lang
	if !IsValidLanguage(lang) {
		http.Error(w, "Unsupported language", http.StatusBadRequest)
		return
	}

	text := reqBody.Text

	value := google_translate.Translate{
		Text: text,
		To:   lang,
	}

	translated, err := google_translate.Translator(value)
	if err != nil {
		panic(err)
	} else {
		type TranslationResponse struct {
			Text string `json:"text"`
		}

		var response TranslationResponse
		prettyJSON, err := json.MarshalIndent(translated, "", "\t")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(prettyJSON, &response)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Default().Printf("Something went wrong while writing response: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
