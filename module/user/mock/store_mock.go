package mock

import (
	"context"
	"quizen/module/user/model"
	"quizen/module/user/store"
)

type UserStoreMock struct {
	WithTransactionFn   func(ctx context.Context, txFunc func(ctx context.Context, tx store.Store) error) error
	CreateUserFn        func(ctx context.Context, user *model.User) (*model.User, error)
	GetUserFn           func(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.User, error)
	UpdateUserFn        func(ctx context.Context, conditions map[string]interface{}, user *model.User) (*model.User, error)
	CreateVerifyEmailFn func(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error)
	UpdateVerifyEmailFn func(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error)
}

func (m UserStoreMock) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, tx store.Store) error) error {
	return m.WithTransactionFn(ctx, txFunc)
}

func (m UserStoreMock) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return m.CreateUserFn(ctx, user)
}

func (m UserStoreMock) GetUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.User, error) {
	return m.GetUserFn(ctx, conditions, moreInfos...)
}

func (m UserStoreMock) UpdateUser(ctx context.Context, conditions map[string]interface{}, user *model.User) (*model.User, error) {
	return m.UpdateUserFn(ctx, conditions, user)
}

func (m UserStoreMock) CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error) {
	return m.CreateVerifyEmailFn(ctx, verifyEmail)
}

func (m UserStoreMock) UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
	return m.UpdateVerifyEmailFn(ctx, email, secretCode)
}
