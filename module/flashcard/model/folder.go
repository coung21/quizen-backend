package model

import (
	"quizen/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	common.SQLModel
	UserID     uuid.UUID `json:"user_id" gorm:"column:user_id"`
	SetID      uuid.UUID `json:"set_id" gorm:"column:set_id"`
	FolderName string    `json:"folder_name" gorm:"column:folder_name"`
}

func (Folder) TableName() string {
	return "folders"
}

func (f *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	f.ID = uuid.New()
	return
}
