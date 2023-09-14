package domain

import (
	"context"
	"errors"

	middl "github.com/diegosperes/pismo/app/middleware"
	"github.com/diegosperes/pismo/app/model"
	"gorm.io/gorm"
)

var (
	ErrTransactionAmountNegative = errors.New("transaction amount must be negative for the given operation type")
	ErrTransactionAmountPositive = errors.New("transaction amount must be positive for the given operation type")
)

func CreateTransaction(ctx context.Context, t *model.Transaction) error {
	if t.OperationTypeId == model.OperationTypePayment && t.Amount < 0 {
		return ErrTransactionAmountPositive
	}

	if t.OperationTypeId != model.OperationTypePayment && t.Amount > 0 {
		return ErrTransactionAmountNegative
	}

	db := ctx.Value(middl.DatabaseContextKey{}).(gorm.DB)
	return db.WithContext(ctx).Create(t).Error
}
