package domain

import (
	"log/slog"

	"github.com/google/uuid"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateAccount(a *model.Account) error {
	return util.GetDatabase().Create(a).Error
}

func GetAccount(id uuid.UUID) (*model.Account, error) {
	a := &model.Account{}
	slog.Info("GetAccount", "id", id)
	db := util.GetDatabase().First(a, id)
	return a, db.Error
}
