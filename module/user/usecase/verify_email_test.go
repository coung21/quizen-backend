package usecase

import (
	"context"
	"quizen/component/token"
	"quizen/module/user/mock"
	"quizen/module/user/model"
	"quizen/module/user/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyEmail(t *testing.T) {
	//Mocking
	var userStoreMock mock.UserStoreMock

	userStoreMock.WithTransactionFn = func(ctx context.Context, txFunc func(ctx context.Context, tx store.Store) error) error {
		return txFunc(ctx, &userStoreMock)
	}

	userStoreMock.UpdateUserFn = func(ctx context.Context, conditions map[string]interface{}, user *model.User) (*model.User, error) {
		return &model.User{
			Email:     user.Email,
			IsVerifed: true,
		}, nil
	}

	userStoreMock.UpdateVerifyEmailFn = func(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
		return &model.VerifyEmail{
			Email:  email,
			IsUsed: true,
		}, nil
	}

	tokenProvider := token.NewJWTProvider("aaaa", 1, 1)

	//Test
	userUsecase := NewUserUsecase(&userStoreMock, nil, tokenProvider)

	t.Run("Success", func(t *testing.T) {
		email := "user1@gg.com"
		code := "123456"
		got, err := userUsecase.VerifyEmail(context.Background(), email, code)

		assert.Nilf(t, err, "error should be nil")
		assert.Equalf(t, email, got.Email, "email should be equal")
		assert.Truef(t, got.IsUsed, "isUsed should be true")
	})

}
