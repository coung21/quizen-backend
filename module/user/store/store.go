package store

import (
	"context"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type Store interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}
