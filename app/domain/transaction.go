package domain

import (
	"context"
	"errors"

	"github.com/diegosperes/pismo/app/model"
	"gorm.io/gorm"
)

type TransactionDatabaseContextKey struct{}

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

	ctxDb := ctx.Value(TransactionDatabaseContextKey{}).(*gorm.DB)
	return ctxDb.WithContext(ctx).Create(t).Error
}
