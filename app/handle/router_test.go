package handle

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/diegosperes/pismo/app/util"
)

func TestServeNotFound(t *testing.T) {
	router := GetConfiguredRouter(&util.AppDependencies{})

	request, _ := http.NewRequest(http.MethodGet, "/not/found/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &util.JsonResponseError{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, responseData.Errors[0].Reason, http.StatusText(http.StatusNotFound))
	assert.Equal(t, responseData.Errors[0].Message, "")
}

func TestServeMethodNotAllowed(t *testing.T) {
	router := GetConfiguredRouter(&util.AppDependencies{})

	request, _ := http.NewRequest("UNKNOW", "/accounts/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &util.JsonResponseError{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
	assert.Equal(t, responseData.Errors[0].Reason, http.StatusText(http.StatusMethodNotAllowed))
	assert.Equal(t, responseData.Errors[0].Message, "")
}
