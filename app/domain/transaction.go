package domain

import (
	"github.com/diegosperes/pismo/app/model"
	"github.com/diegosperes/pismo/app/util"
)

func CreateTransaction(t *model.Transaction) error {
	return util.GetDatabase().Create(t).Error
}
