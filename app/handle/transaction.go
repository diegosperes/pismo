package handle

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/diegosperes/pismo/app/domain"
	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	transaction := &model.Transaction{}

	if jsonErr := json.NewDecoder(r.Body).Decode(&transaction); jsonErr != nil {
		util.WriteJsonError(w, http.StatusBadRequest, jsonErr.Error())
		return
	}

	ctx := r.Context()
	ctx = context.WithValue(ctx, domain.TransactionDatabaseContextKey{}, util.GetDatabase())
	transactionErr := domain.CreateTransaction(ctx, transaction)

	if transactionErr != nil {
		util.WriteJsonError(w, http.StatusBadRequest, transactionErr.Error())
		return
	}

	util.WriteJson(w, transaction, http.StatusCreated)
}
