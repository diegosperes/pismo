package domain

import (
	"context"

	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateTransaction(ctx context.Context, t *model.Transaction) error {
	return util.GetDatabase().WithContext(ctx).Create(t).Error
}
