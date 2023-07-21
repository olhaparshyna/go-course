package middleware

import (
	"github.com/joho/godotenv"
	"go-course/home_assignment14/config"
	"net/http"
)

func IsValidApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			http.Error(w, "Error reading .env file", http.StatusInternalServerError)
			return
		}

		conf := config.New()

		validApiKey := conf.ApiKey

		apiKey := r.Header.Get("API-KEY")

		if apiKey != validApiKey {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
