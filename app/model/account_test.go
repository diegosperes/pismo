package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountWithouDocumentNumber(t *testing.T) {
	account := &Account{}
	assert.ErrorIs(t, account.BeforeCreate(nil), ErrAccountDocumentNumberRequired)
}
