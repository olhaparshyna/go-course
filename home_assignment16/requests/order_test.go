package requests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCreateRequestData(t *testing.T) {
	validJSON := `{
		"name": "John",
		"email": "john@example.com",
		"items": ["item1", "item2"]
	}`
	reqValid := httptest.NewRequest("POST", "/", bytes.NewBufferString(validJSON))

	recorder := httptest.NewRecorder()

	data := ValidateCreateRequestData(recorder, reqValid)
	assert.NotNil(t, data)
	assert.Equal(t, "John", data.Name)
	assert.Equal(t, "john@example.com", data.Email)
	assert.Len(t, data.Items, 2)
}

func TestInvalidateCreateRequestData(t *testing.T) {
	missingFieldsJSON := `{
		"name": "John",
		"email": ""
	}`
	reqMissingFields := httptest.NewRequest("POST", "/", bytes.NewBufferString(missingFieldsJSON))

	recorder := httptest.NewRecorder()

	ValidateCreateRequestData(recorder, reqMissingFields)
	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}
