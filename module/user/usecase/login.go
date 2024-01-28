package usecase

import (
	"context"
	"quizen/common"
	"quizen/module/user/model"
)

type tokensResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u userUsecase) Login(ctx context.Context, email, password string) (*model.User, *tokensResp, error) {
	foundUser, err := u.userStore.GetUser(ctx, map[string]interface{}{"email": email, "is_verified": true})
	if foundUser == nil && err != nil {
		return nil, nil, common.NotExistAccount
	}

	if err := foundUser.ComparePassword(password); err != nil {
		return nil, nil, common.WrongPassword
	}

	tokenPayload := u.tokenProvider.NewPayLoad(foundUser.ID)

	accessToken, refreshToken, err := u.tokenProvider.GenerateTokens(tokenPayload)

	if err != nil {
		return nil, nil, err
	}

	return foundUser, &tokensResp{AccessToken: *accessToken, RefreshToken: *refreshToken}, nil

}
