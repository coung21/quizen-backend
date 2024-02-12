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

func TestRenewToken(t *testing.T) {
	mockUserStore := mock.UserStoreMock{}

	findingSessionID := uuid.New()
	findingUserID := uuid.New()

	tokenProvider := token.NewJWTProvider("aaaa", 1, 1)
	payload := tokenProvider.NewPayLoad(findingUserID)
	_, findingRefreshToken, _ := tokenProvider.GenerateTokens(payload)

	mockUserStore.GetSessionFn = func(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.Session, error) {
		// list of sessions
		sessions := []model.Session{
			{ID: findingSessionID, UserID: findingUserID, RefreshToken: *findingRefreshToken, ExpiresAt: time.Now().Add(time.Hour)},
			{ID: uuid.New(), UserID: uuid.New(), RefreshToken: "abc", ExpiresAt: time.Now().Add(time.Hour)},
		}

		for _, session := range sessions {
			if session.ID.String() == conditions["id"] {
				return &session, nil
			}

		}

		return nil, common.NotFound
	}

	userUsecase := NewUserUsecase(&mockUserStore, nil, tokenProvider)

	t.Run("Valid case", func(t *testing.T) {
		// Setup
		sessionID := findingSessionID.String()
		refreshToken := *findingRefreshToken

		// Execution
		accessToken, err := userUsecase.RenewToken(context.Background(), sessionID, refreshToken)

		// Validation
		assert.Nilf(t, err, "Error should be nil")
		assert.NotEmptyf(t, accessToken, "Access token should not be empty")
	})

	t.Run("Invalid session id", func(t *testing.T) {
		// Setup
		sessionID := uuid.New().String()
		refreshToken := *findingRefreshToken

		// Execution
		accessToken, err := userUsecase.RenewToken(context.Background(), sessionID, refreshToken)

		// Validation
		assert.Equalf(t, common.NotFound, err, "Error should be not found")
		assert.Emptyf(t, accessToken, "Access token should be empty")
	})

	t.Run("Invalid refresh token", func(t *testing.T) {
		// Setup
		sessionID := findingSessionID.String()
		refreshToken := "abc"

		// Execution
		accessToken, err := userUsecase.RenewToken(context.Background(), sessionID, refreshToken)

		// Validation
		assert.Errorf(t, err, "Should return error")
		assert.Emptyf(t, accessToken, "Access token should be empty")
	})
}
