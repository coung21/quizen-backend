package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"session_id" gorm:"column:id"`
	UserID       string    `json:"user_id" gorm:"column:user_id"`
	RefreshToken string    `json:"refresh_token" gorm:"column:refresh_token"`
	UserAgent    string    `json:"user_agent" gorm:"column:user_agent"`
	UserIP       string    `json:"user_ip" gorm:"column:user_ip"`
	ExpriesAt    time.Time `json:"expires_at" gorm:"column:expires_at"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Session) TableName() string { return "sessions" }

func (s *Session) BeforeCreate() error {
	s.ID = uuid.New()
	return nil
}
