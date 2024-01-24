package store

import (
	"context"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type Store interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUserByEmail(context.Context, string) (*model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	UpdateUser(context.Context, int, *model.User) (*model.User, error)
	CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}
