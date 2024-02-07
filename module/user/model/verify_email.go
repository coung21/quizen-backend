package model

import (
	"quizen/common"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerifyEmail struct {
	common.SQLModel
	Email      string    `json:"email" gorm:"column:email"`
	SecretCode string    `json:"secret_code,omitempty" gorm:"column:secret_code"`
	IsUsed     bool      `json:"is_used" gorm:"column:is_used"`
	ExpriedAt  time.Time `json:"expired_at,omitempty" gorm:"column:expired_at;default:(current_timestamp + interval 10 minute)"`
}

func (VerifyEmail) TableName() string {
	return "verify_emails"
}

func (e *VerifyEmail) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.New()
	return nil
}
