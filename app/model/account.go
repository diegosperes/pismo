package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrAccountDocumentNumberRequired = errors.New("Account document number is required")
)

type Account struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	DocumentNumber string         `json:"document_number" gorm:"unique;not null"`
	CreatedAt      time.Time      `json:"event_date" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if len(a.DocumentNumber) == 0 {
		return ErrAccountDocumentNumberRequired
	}

	return nil
}
