package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func init() {
	util.SetupApp()
}

func newTestAccount() *model.Account {
	return &model.Account{
		DocumentNumber: strconv.Itoa(rand.Intn(99999999999)),
	}
}

func TestCreateAccount(t *testing.T) {
	router := GetConfiguredRouter()

	createdAccount := newTestAccount()
	jsonBody, _ := json.Marshal(createdAccount)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/accounts/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.NotEqual(t, uuid.Nil, responseData.ID)
}

func TestCreateInvalidAccount(t *testing.T) {
	router := GetConfiguredRouter()

	createdAccount := &model.Account{}
	jsonBody, _ := json.Marshal(createdAccount)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/accounts/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestGetAccount(t *testing.T) {
	router := GetConfiguredRouter()

	createdAccount := newTestAccount()
	util.GetDatabase().Create(createdAccount)

	requestPath := fmt.Sprintf("/accounts/%s", createdAccount.ID.String())
	request, _ := http.NewRequest(http.MethodGet, requestPath, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, createdAccount.ID, responseData.ID)
	assert.Equal(t, createdAccount.DocumentNumber, responseData.DocumentNumber)
}

func TestGetNonExistingAccount(t *testing.T) {
	router := GetConfiguredRouter()

	requestPath := fmt.Sprintf("/accounts/%s", uuid.New().String())
	request, _ := http.NewRequest(http.MethodGet, requestPath, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusNotFound, response.Code)
}
