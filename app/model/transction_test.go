package model

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestTransaction(opt OperationType, amount float64) *Transaction {
	return &Transaction{
		AccountId:       uuid.New(),
		OperationTypeId: opt,
		Amount:          amount,
	}
}

func TestTransactionWithOperationTypeLumpSum(t *testing.T) {
	transaction := newTestTransaction(OperationTypeLumpSum, 100.0)
	assert.ErrorIs(t, transaction.BeforeCreate(nil), ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypeInstallments(t *testing.T) {
	transaction := newTestTransaction(OperationTypeInstallments, 100.0)
	assert.ErrorIs(t, transaction.BeforeCreate(nil), ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypeWithdraw(t *testing.T) {
	transaction := newTestTransaction(OperationTypeWithdraw, 100.0)
	assert.ErrorIs(t, transaction.BeforeCreate(nil), ErrTransactionAmountNegative)
}

func TestTransactionWithOperationTypePayment(t *testing.T) {
	transaction := newTestTransaction(OperationTypePayment, -100.0)
	assert.ErrorIs(t, transaction.BeforeCreate(nil), ErrTransactionAmountPositive)
}
