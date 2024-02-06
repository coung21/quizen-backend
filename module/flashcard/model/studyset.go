package model

import (
	"quizen/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudySet struct {
	common.SQLModel
	UserID      uuid.UUID `json:"user_id" gorm:"column:user_id"`
	SetName     string    `json:"set_name" binding:"required" gorm:"column:set_name"`
	Description string    `json:"description" gorm:"column:description"`
}

func (StudySet) TableName() string {
	return "study_sets"
}

func (s *StudySet) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.New()
	return
}
