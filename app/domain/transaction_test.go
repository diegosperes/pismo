package domain

import (
	"context"
	"testing"

	"github.com/diegosperes/pismo/app/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestTransaction(opt model.OperationType, amount float64) *model.Transaction {
	return &model.Transaction{
		AccountId:       uuid.New(),
		OperationTypeId: opt,
		Amount:          amount,
	}
}

func TestTransactionWithOperationTypeLumpSum(t *testing.T) {
	transaction := newTestTransaction(model.OperationTypeLumpSum, 100.0)
	err := CreateTransaction(context.Background(), transaction)
	assert.ErrorIs(t, err, ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypeInstallments(t *testing.T) {
	transaction := newTestTransaction(model.OperationTypeInstallments, 100.0)
	err := CreateTransaction(context.Background(), transaction)
	assert.ErrorIs(t, err, ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypeWithdraw(t *testing.T) {
	transaction := newTestTransaction(model.OperationTypeWithdraw, 100.0)
	err := CreateTransaction(context.Background(), transaction)
	assert.ErrorIs(t, err, ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypePayment(t *testing.T) {
	transaction := newTestTransaction(model.OperationTypePayment, -100.0)
	err := CreateTransaction(context.Background(), transaction)
	assert.ErrorIs(t, err, ErrTransactionAmountPositive)
}
