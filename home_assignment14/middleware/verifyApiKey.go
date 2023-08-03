package middleware

import (
	"go-course/home_assignment14/config"
	"net/http"
)

func IsValidApiKey(conf *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			validApiKey := conf.ApiKey

			apiKey := r.Header.Get("API-KEY")

			if apiKey != validApiKey {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
