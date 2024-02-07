package store

import (
	"context"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type Store interface {
	WithTransaction(ctx context.Context, txFunc func(ctx context.Context, tx Store) error) error
	CreateUser(context.Context, *model.User) (*model.User, error)
	GetUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.User, error)
	UpdateUser(ctx context.Context, conditions map[string]interface{}, user *model.User) (*model.User, error)
	CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error)
	UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error)
	CreateSession(ctx context.Context, session *model.Session) (*model.Session, error)
	GetSession(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.Session, error)
	DeleteSession(ctx context.Context, conditions map[string]interface{}) error
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return UserStore{db: db}
}

func (s UserStore) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, tx Store) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txUserStore := UserStore{db: tx}

	err := txFunc(ctx, txUserStore)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
