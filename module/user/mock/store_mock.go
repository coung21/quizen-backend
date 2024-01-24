package mock

import (
	"context"
	"quizen/module/user/model"
)

type UserStoreMock struct {
	CreateUserFn        func(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByEmailFn    func(ctx context.Context, email string) (*model.User, error)
	GetUserByIdFn       func(ctx context.Context, id int) (*model.User, error)
	UpdateUserFn        func(ctx context.Context, id int, user *model.User) (*model.User, error)
	CreateVerifyEmailFn func(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error)
	UpdateVerifyEmailFn func(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error)
}

func (m *UserStoreMock) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return m.CreateUserFn(ctx, user)
}

func (m *UserStoreMock) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.GetUserByEmailFn(ctx, email)
}

func (m *UserStoreMock) GetUserById(ctx context.Context, id int) (*model.User, error) {
	return m.GetUserByIdFn(ctx, id)
}

func (m *UserStoreMock) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	return m.UpdateUserFn(ctx, id, user)
}

func (m *UserStoreMock) CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error) {
	return m.CreateVerifyEmailFn(ctx, verifyEmail)
}

func (m *UserStoreMock) UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
	return m.UpdateVerifyEmailFn(ctx, email, secretCode)
}
