package handle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func newTestTransaction(db gorm.DB) *model.Transaction {
	account := newTestAccount()
	db.Create(account)

	return &model.Transaction{
		AccountId:       account.ID,
		OperationTypeId: model.OperationTypePayment,
		Amount:          100.00,
	}
}

type TransactionTestSuite struct {
	suite.Suite

	dependencies *util.AppDependencies
	router       http.Handler
}

func (s *TransactionTestSuite) SetupSuite() {
	s.dependencies = util.SetupApp()
	s.router = GetConfiguredRouter(s.dependencies)
}

func (s *TransactionTestSuite) TestCreateTransaction() {
	createdTransaction := newTestTransaction(s.dependencies.Database)
	jsonBody, _ := json.Marshal(createdTransaction)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/transactions/", requestBody)
	response := httptest.NewRecorder()

	s.router.ServeHTTP(response, request)

	responseData := &model.Transaction{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	s.Equal(http.StatusCreated, response.Code)
	s.NotEqual(uuid.Nil, responseData.ID)
	s.Equal(createdTransaction.AccountId, responseData.AccountId)
	s.Equal(createdTransaction.OperationTypeId, responseData.OperationTypeId)
	s.Equal(createdTransaction.Amount, responseData.Amount)
}

func (s *TransactionTestSuite) TestCreateInvalidTransaction() {
	createdTransaction := newTestTransaction(s.dependencies.Database)
	createdTransaction.OperationTypeId = model.OperationTypeLumpSum
	jsonBody, _ := json.Marshal(createdTransaction)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/transactions/", requestBody)
	response := httptest.NewRecorder()

	s.router.ServeHTTP(response, request)

	s.Equal(http.StatusBadRequest, response.Code)
}
