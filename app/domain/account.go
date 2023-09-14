package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/diegosperes/pismo/app/model"
)

type AccountDatabaseContextKey struct{}

var (
	ErrAccountDocumentNumberRequired = errors.New("account document number is required")
)

func CreateAccount(ctx context.Context, a *model.Account) error {
	if len(a.DocumentNumber) == 0 {
		return ErrAccountDocumentNumberRequired
	}

	ctxDb := ctx.Value(AccountDatabaseContextKey{}).(*gorm.DB)
	return ctxDb.WithContext(ctx).Create(a).Error
}

func GetAccount(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	a := &model.Account{}
	ctxDb := ctx.Value(AccountDatabaseContextKey{}).(*gorm.DB)
	resultDb := ctxDb.WithContext(ctx).First(a, id)
	return a, resultDb.Error
}
