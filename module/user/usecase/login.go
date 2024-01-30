package usecase

import (
	"context"
	"quizen/common"
	"quizen/module/user/model"
	"time"
)

type tokensResp struct {
	AccessToken     string    `json:"access_token"`
	RefreshToken    string    `json:"refresh_token"`
	RefreshTokenExp time.Time `json:"refresh_token_exp"`
}

func (u userUsecase) Login(ctx context.Context, email, password string) (*model.User, *tokensResp, string, error) {
	foundUser, err := u.userStore.GetUser(ctx, map[string]interface{}{"email": email, "is_verified": true})
	if foundUser == nil && err != nil {
		return nil, nil, "", common.NotExistAccount
	}

	if err := foundUser.ComparePassword(password); err != nil {
		return nil, nil, "", common.WrongPassword
	}

	tokenPayload := u.tokenProvider.NewPayLoad(foundUser.ID)

	accessToken, refreshToken, err := u.tokenProvider.GenerateTokens(tokenPayload)

	session, err := u.userStore.CreateSession(ctx, &model.Session{
		UserID:       foundUser.ID,
		RefreshToken: *refreshToken,
		UserAgent:    ctx.Value("user_agent").(string),
		UserIP:       ctx.Value("user_ip").(string),
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 7),
	})

	if err != nil {
		return nil, nil, "", err
	}

	return foundUser, &tokensResp{AccessToken: *accessToken, RefreshToken: *refreshToken, RefreshTokenExp: session.ExpiresAt}, session.ID.String(), nil

}
