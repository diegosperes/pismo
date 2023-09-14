package domain

import (
	"context"
	"testing"

	"github.com/diegosperes/pismo/app/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountWithouDocumentNumber(t *testing.T) {
	account := &model.Account{}
	err := CreateAccount(context.Background(), account)
	assert.ErrorIs(t, err, ErrAccountDocumentNumberRequired)
}
