package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID              uuid.UUID      `json:"id" gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	AccountId       uuid.UUID      `json:"account_id" gorm:"type:uuid;not null"`
	OperationTypeId OperationType  `json:"operation_type_id" gorm:"not null"`
	Amount          float64        `json:"amount" gorm:"not null"`
	CreatedAt       time.Time      `json:"event_date" gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
