package common

import "time"

type SQLModel struct {
	ID        int       `json:"id" gorm:"column:id"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;"`
}
