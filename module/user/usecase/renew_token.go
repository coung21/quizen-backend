package usecase

import (
	"context"
	"quizen/common"
	"time"
)

func (u userUsecase) RenewToken(ctx context.Context, sessionID string, refreshToken string) (string, error) {
	//validate session
	session, err := u.userStore.GetSession(ctx, map[string]interface{}{"id": sessionID})
	if err != nil {
		return "", err
	}
	//validate token
	tokenClaims, err := u.tokenProvider.Validate(refreshToken)
	if err != nil {
		return "", err
	}

	if tokenClaims.ID != session.UserID {
		return "", common.Unauthorized
	}

	if refreshToken != session.RefreshToken {
		return "", common.Forbidden
	}

	//validate session and token
	if session.ExpiresAt.Before(time.Now()) {
		return "", common.ErrJWTExpired
	}

	//generate new token
	tokenPayload := u.tokenProvider.NewPayLoad(session.UserID)
	access, _, err := u.tokenProvider.GenerateTokens(tokenPayload)

	return *access, nil
}
