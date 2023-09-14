package domain

import (
	"context"

	"github.com/google/uuid"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateAccount(ctx context.Context, a *model.Account) error {
	return util.GetDatabase().WithContext(ctx).Create(a).Error
}

func GetAccount(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	a := &model.Account{}
	db := util.GetDatabase().WithContext(ctx).First(a, id)
	return a, db.Error
}
