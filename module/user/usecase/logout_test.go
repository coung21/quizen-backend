package usecase

import (
	"context"
	"quizen/common"
	"quizen/module/user/mock"
	"quizen/module/user/model"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestLogout(t *testing.T) {
	userStoreMock := mock.UserStoreMock{}

	logoutSessionID := uuid.New()

	userStoreMock.DeleteSessionFn = func(ctx context.Context, conditions map[string]interface{}) error {
		sessionsArr := []model.Session{
			{ID: logoutSessionID,
				UserID: uuid.New(),
			},
		}

		for _, session := range sessionsArr {
			if session.ID.String() == conditions["id"].(string) {
				return nil
			}

		}
		return common.NotFound
	}

	userUsecase := NewUserUsecase(userStoreMock, nil, nil)

	t.Run("Logout success", func(t *testing.T) {
		err := userUsecase.Logout(context.Background(), logoutSessionID.String())

		assert.Nil(t, err)
	})

	t.Run("Logout failed", func(t *testing.T) {
		err := userUsecase.Logout(context.Background(), uuid.New().String())

		assert.NotNil(t, err)
	})
}
