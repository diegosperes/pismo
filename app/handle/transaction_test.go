package handle

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

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

func TestCreateTransaction(t *testing.T) {
	router := GetConfiguredRouter()

	createdTransaction := newTestTransaction()
	jsonBody, _ := json.Marshal(createdTransaction)

	requestBody := bytes.NewReader(jsonBody)
	request, _ := http.NewRequest(http.MethodPost, "/transactions/", requestBody)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	responseData := &model.Transaction{}
	json.Unmarshal(response.Body.Bytes(), responseData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.NotEqual(t, uuid.Nil, responseData.ID)
	assert.Equal(t, createdTransaction.AccountId, responseData.AccountId)
	assert.Equal(t, createdTransaction.OperationTypeId, responseData.OperationTypeId)
	assert.Equal(t, createdTransaction.Amount, responseData.Amount)
}
