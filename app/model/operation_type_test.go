package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperationTypeLumpSum(t *testing.T) {
	assert.Equal(t, OperationTypeLumpSum.GetName(), "LumpSum")
}

func TestOperationTypeInstallments(t *testing.T) {
	assert.Equal(t, OperationTypeInstallments.GetName(), "Installments")
}

func TestOperationTypeWithdraw(t *testing.T) {
	assert.Equal(t, OperationTypeWithdraw.GetName(), "Withdraw")
}

func TestOperationTypePayment(t *testing.T) {
	assert.Equal(t, OperationTypePayment.GetName(), "Payment")
}
