package handle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func newTestAccount() *model.Account {
	return &model.Account{
		DocumentNumber: strconv.Itoa(rand.Intn(99999999999)),
	}
}

type AccountTestSuite struct {
	suite.Suite
}

func (s *AccountTestSuite) SetupSuite() {
	util.SetupApp()
}

func (s *AccountTestSuite) TestCreateAccount() {
	router := GetConfiguredRouter()

	createdAccount := newTestAccount()
	jsonBody, _ := json.Marshal(createdAccount)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/accounts/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusCreated, response.Code)
	s.NotEqual(uuid.Nil, responseData.ID)
}

func (s *AccountTestSuite) TestCreateInvalidAccount() {
	router := GetConfiguredRouter()

	createdAccount := &model.Account{}
	jsonBody, _ := json.Marshal(createdAccount)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/accounts/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusBadRequest, response.Code)
}

func (s *AccountTestSuite) TestGetAccount() {
	router := GetConfiguredRouter()

	createdAccount := newTestAccount()
	util.GetDatabase().Create(createdAccount)

	requestPath := fmt.Sprintf("/accounts/%s", createdAccount.ID.String())
	request, _ := http.NewRequest(http.MethodGet, requestPath, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusOK, response.Code)
	s.Equal(createdAccount.ID, responseData.ID)
	s.Equal(createdAccount.DocumentNumber, responseData.DocumentNumber)
}

func (s *AccountTestSuite) TestGetNonExistingAccount() {
	router := GetConfiguredRouter()

	requestPath := fmt.Sprintf("/accounts/%s", uuid.New().String())
	request, _ := http.NewRequest(http.MethodGet, requestPath, nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Account{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusNotFound, response.Code)
}
