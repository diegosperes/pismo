package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	middl "github.com/diegosperes/pismo/app/middleware"
	"github.com/diegosperes/pismo/app/model"
)

var (
	ErrAccountDocumentNumberRequired = errors.New("account document number is required")
)

func CreateAccount(ctx context.Context, a *model.Account) error {
	if len(a.DocumentNumber) == 0 {
		return ErrAccountDocumentNumberRequired
	}

	db := ctx.Value(middl.DatabaseContextKey{}).(gorm.DB)
	return db.WithContext(ctx).Create(a).Error
}

func GetAccount(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	a := &model.Account{}
	db := ctx.Value(middl.DatabaseContextKey{}).(gorm.DB)
	resultDb := db.WithContext(ctx).First(a, id)
	return a, resultDb.Error
}
