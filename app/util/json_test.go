package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type jsonTestResponse struct {
	Name string `json:"name"`
}

func TestWriteJsonResponse(t *testing.T) {
	response := httptest.NewRecorder()
	WriteJson(response, &jsonTestResponse{Name: "Diego"}, http.StatusOK)

	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Equal(t, "{\"name\":\"Diego\"}\n", response.Body.String())
}

func TestWriteJsonResponseError(t *testing.T) {
	response := httptest.NewRecorder()
	WriteJsonError(response, http.StatusNotFound, "Something not found")
	expected := "{\"errors\":[{\"reason\":\"Not Found\",\"message\":\"Something not found\"}]}\n"

	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	assert.Equal(t, expected, response.Body.String())
}
