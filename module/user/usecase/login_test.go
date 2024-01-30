package usecase

import (
	"context"
	"quizen/common"
	"quizen/component/token"
	"quizen/module/user/mock"
	"quizen/module/user/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// Setup
	now := time.Now()
	// Mocking
	userStoreMock := mock.UserStoreMock{}

	userStoreMock.GetUserFn = func(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.User, error) {
		var result *model.User

		userArr := []model.User{
			{SQLModel: common.SQLModel{ID: uuid.New(), CreatedAt: &now, UpdatedAt: &now}, Username: "user1", Email: "user1@gg.com", Password: "password"},
			{SQLModel: common.SQLModel{ID: uuid.New(), CreatedAt: &now, UpdatedAt: &now}, Username: "user2", Email: "user2@gg.com", Password: "password"},
		}

		for _, user := range userArr {
			if user.Email == conditions["email"] {
				result = &user
				break
			}
		}

		if result == nil {
			return nil, common.NotFound
		}

		result.HashPassword()

		return result, nil
	}

	userStoreMock.CreateSessionFn = func(ctx context.Context, session *model.Session) (*model.Session, error) {

		return &model.Session{
			ID:           uuid.New(),
			UserID:       session.UserID,
			RefreshToken: session.RefreshToken,
			UserAgent:    session.UserAgent,
			UserIP:       session.UserIP,
			ExpiresAt:    session.ExpiresAt,
			CreatedAt:    now,
			UpdatedAt:    now,
		}, nil
	}

	tokenProvider := token.NewJWTProvider("aaaa", 1, 1)

	//inject mock
	userUsecase := NewUserUsecase(&userStoreMock, nil, tokenProvider)

	t.Run("Success", func(t *testing.T) {
		email := "user1@gg.com"
		password := "password"
		gotUser, gotTokens, gotSessionID, err := userUsecase.Login(context.Background(), email, password)

		assert.Nilf(t, err, "error should be nil")
		assert.NotEmptyf(t, gotSessionID, "sessionID should not be empty")
		assert.Equalf(t, email, gotUser.Email, "email should be equal")
		assert.Equalf(t, "user1", gotUser.Username, "username should be equal")
		assert.NotEmptyf(t, gotTokens.AccessToken, "accessToken should not be empty")
		assert.NotEmptyf(t, gotTokens.RefreshToken, "refreshToken should not be empty")
	})

	t.Run("Wrong password", func(t *testing.T) {
		email := "user1@gg.com"
		password := "wrongpassword"
		gotUser, gotTokens, gotSessionID, err := userUsecase.Login(context.Background(), email, password)

		assert.Equalf(t, common.WrongPassword, err, "error should be WrongPassword")
		assert.Emptyf(t, gotSessionID, "sessionID should be empty")
		assert.Nilf(t, gotUser, "user should be nil")
		assert.Nilf(t, gotTokens, "tokens should be nil")
	})

	t.Run("Not exist account", func(t *testing.T) {
		email := "user3@gg.cpm"
		password := "password"
		gotUser, gotTokens, gotSessionID, err := userUsecase.Login(context.Background(), email, password)

		assert.Equalf(t, common.NotExistAccount, err, "error should be NotExistAccount")
		assert.Emptyf(t, gotSessionID, "sessionID should be empty")
		assert.Nilf(t, gotUser, "user should be nil")
		assert.Nilf(t, gotTokens, "tokens should be nil")
	})
}
