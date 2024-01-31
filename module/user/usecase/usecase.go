package usecase

import (
	"context"
	"quizen/component/token"
	"quizen/component/worker"
	"quizen/module/user/model"
	"quizen/module/user/store"
)

type Usecase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	VerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error)
	Login(ctx context.Context, email, password string) (*model.User, *tokensResp, string, error)
	Logout(ctx context.Context, sessionID string) error
	RenewToken(ctx context.Context, sessionID string, refreshToken string) (string, error)
}

type userUsecase struct {
	userStore      store.Store
	taskDistrbutor worker.TaskDistributor
	tokenProvider  token.TokenProvider
}

func NewUserUsecase(userStore store.Store, taskDistributor worker.TaskDistributor, tokenProvider token.TokenProvider) userUsecase {
	return userUsecase{userStore: userStore, taskDistrbutor: taskDistributor, tokenProvider: tokenProvider}
}
