package handle

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"

	"github.com/diegosperes/pismo/app/domain"
	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account := &model.Account{}

	if jsonErr := json.NewDecoder(r.Body).Decode(&account); jsonErr != nil {
		util.WriteJsonError(w, http.StatusBadRequest, jsonErr.Error())
		return
	}

	accountErr := domain.CreateAccount(r.Context(), account)

	if accountErr != nil {
		util.WriteJsonError(w, http.StatusBadRequest, accountErr.Error())
		return
	}

	util.WriteJson(w, account, http.StatusCreated)
}

func GetAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	accountId, uuidErr := uuid.Parse(ps.ByName("accountId"))

	if uuidErr != nil {
		util.WriteJsonError(w, http.StatusBadRequest, uuidErr.Error())
		return
	}

	account, accountErr := domain.GetAccount(r.Context(), accountId)

	if errors.Is(accountErr, gorm.ErrRecordNotFound) {
		util.WriteJsonError(w, http.StatusNotFound)
		return
	}

	if accountErr != nil {
		slog.Error("An error ocurred on GetAccount handle", "accountId", accountId, "error", accountErr.Error())
		util.WriteJsonError(w, http.StatusInternalServerError)
		return
	}

	util.WriteJson(w, account, http.StatusOK)
}
