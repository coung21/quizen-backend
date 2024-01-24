package store

import (
	"context"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type Store interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUser(context.Context, int) (*model.User, error)
	GetUserByEmail(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, int, *model.User) (*model.User, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}
