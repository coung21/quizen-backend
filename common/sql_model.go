package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
}
