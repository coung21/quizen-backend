package model

import (
	"quizen/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flashcard struct {
	common.SQLModel
	StudySetID uuid.UUID     `json:"study_set_id" gorm:"column:study_set_id"`
	Term       string        `json:"term" binding:"required" gorm:"column:term"`
	Definition string        `json:"definition" binding:"required" gorm:"column:definition"`
	Image      *common.Image `json:"image" gorm:"column:image"`
}

func (Flashcard) TableName() string {
	return "flashcards"
}

func (f *Flashcard) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}
