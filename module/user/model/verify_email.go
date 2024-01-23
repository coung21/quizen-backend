package model

import (
	"quizen/common"
	"time"
)

type VerifyEmail struct {
	common.SQLModel
	Email      string    `json:"email" gorm:"column:email" validate:"required,email"`
	SecretCode string    `json:"secret_code" gorm:"column:secret_code" validate:"required"`
	IsUsed     bool      `json:"is_used" gorm:"column:is_used"`
	ExpriedAt  time.Time `json:"expried_at" gorm:"column:expried_at"`
}
