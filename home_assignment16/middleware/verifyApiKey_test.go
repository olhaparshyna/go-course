package middleware

import (
	"github.com/stretchr/testify/assert"
	"go-course/home_assignment16/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidApiKeyMiddleware(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mockConfig := &config.Config{
		ApiKey: "valid-api-key",
	}

	reqValid := httptest.NewRequest("GET", "/", nil)
	reqValid.Header.Set("API-KEY", "valid-api-key")

	reqInvalid := httptest.NewRequest("GET", "/", nil)
	reqInvalid.Header.Set("API-KEY", "invalid-api-key")

	recorder := httptest.NewRecorder()

	middleware := IsValidApiKey(mockConfig)

	handler := middleware(mockHandler)
	handler.ServeHTTP(recorder, reqValid)
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "OK", recorder.Body.String())
}

func TestIsInvalidApiKeyMiddleware(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mockConfig := &config.Config{
		ApiKey: "valid-api-key",
	}

	reqInvalid := httptest.NewRequest("GET", "/", nil)
	reqInvalid.Header.Set("API-KEY", "invalid-api-key")

	recorder := httptest.NewRecorder()

	middleware := IsValidApiKey(mockConfig)

	handler := middleware(mockHandler)
	handler.ServeHTTP(recorder, reqInvalid)
	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "Invalid API Key\n", recorder.Body.String())
}
