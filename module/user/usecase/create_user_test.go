package usecase

import (
	"context"
	"quizen/common"
	"quizen/component/worker"
	"quizen/module/user/mock"
	"quizen/module/user/model"
	"testing"
	"time"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup
	now := time.Now()
	// Mocking
	userStoreMock := mock.UserStoreMock{}
	taskDistributorMock := mock.TaskDistributorMock{}

	userStoreMock.GetUserByEmailFn = func(ctx context.Context, email string) (*model.User, error) {
		var result *model.User

		userArr := []model.User{
			{SQLModel: common.SQLModel{ID: 1, CreatedAt: &now, UpdatedAt: &now}, Username: "user1", Email: "user1@gg.com", Password: "password"},
			{SQLModel: common.SQLModel{ID: 2, CreatedAt: &now, UpdatedAt: &now}, Username: "user2", Email: "user2@gg.com", Password: "password"},
		}

		for _, user := range userArr {
			if user.Email == email {
				result = &user
				break
			}
		}

		if result == nil {
			return nil, common.NotFound
		}
		return result, nil
	}
	userStoreMock.CreateUserFn = func(ctx context.Context, user *model.User) (*model.User, error) {
		return &model.User{
			SQLModel: common.SQLModel{ID: user.ID, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt},
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}, nil
	}

	taskDistributorMock.DistributeTaskSendVerifyEmailFn = func(ctx context.Context, payload *worker.PayloadVerifyEmail, opts ...asynq.Option) error {
		return nil
	}

	// Dependency injection
	userUsecase := NewUserUsecase(&userStoreMock, &taskDistributorMock)

	// Assertion
	t.Run("Create user success", func(t *testing.T) {
		password := "password"
		newUser := &model.User{
			Username: "user3",
			Password: password,
			Email:    "user3@gg.com",
		}
		got, err := userUsecase.CreateUser(context.Background(), newUser)

		// Assert with testify assert package
		assert.Nilf(t, err, "Error should be nil")
		assert.Equal(t, newUser.Username, got.Username, "Username should be equal")
		assert.Equal(t, newUser.Email, got.Email, "Email should be equal")
		assert.NotEqual(t, password, got.Password, "Password should be hashed")
	})

	t.Run("Create user with existing email", func(t *testing.T) {
		newUser := &model.User{
			Username: "user1",
			Password: "password",
			Email:    "user1@gg.com",
		}
		got, err := userUsecase.CreateUser(context.Background(), newUser)

		// Assert with testify assert package
		assert.Nilf(t, got, "User should be nil")
		assert.Equal(t, common.ExistsEmailError, err, "Error should be 'Email already exists'")
	})

}
