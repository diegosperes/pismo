package handle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func newTestTransaction() *model.Transaction {
	account := newTestAccount()
	util.GetDatabase().Create(account)

	return &model.Transaction{
		AccountId:       account.ID,
		OperationTypeId: model.OperationTypePayment,
		Amount:          100.00,
	}
}

type TransactionTestSuite struct {
	suite.Suite
}

func (s *TransactionTestSuite) SetupSuite() {
	util.SetupApp()
}

func (s *TransactionTestSuite) TestCreateTransaction() {
	router := GetConfiguredRouter()

	createdTransaction := newTestTransaction()
	jsonBody, _ := json.Marshal(createdTransaction)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/transactions/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Transaction{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusCreated, response.Code)
	s.NotEqual(uuid.Nil, responseData.ID)
	s.Equal(createdTransaction.AccountId, responseData.AccountId)
	s.Equal(createdTransaction.OperationTypeId, responseData.OperationTypeId)
	s.Equal(createdTransaction.Amount, responseData.Amount)
}

func (s *TransactionTestSuite) TestCreateInvalidTransaction() {
	router := GetConfiguredRouter()

	createdTransaction := newTestTransaction()
	createdTransaction.OperationTypeId = model.OperationTypeLumpSum
	jsonBody, _ := json.Marshal(createdTransaction)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/transactions/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Transaction{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusBadRequest, response.Code)
}
