package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID             uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	DocumentNumber string         `json:"document_number" gorm:"unique;not null"`
	CreatedAt      time.Time      `json:"event_date" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
