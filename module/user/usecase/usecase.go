package usecase

import (
	"context"
	"quizen/component/worker"
	"quizen/module/user/model"
	"quizen/module/user/store"
)

type Usecase interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type userUsecase struct {
	userStore      store.Store
	taskDistrbutor worker.TaskDistributor
}

func NewUserUsecase(userStore store.Store, taskDistributor worker.TaskDistributor) userUsecase {
	return userUsecase{userStore: userStore, taskDistrbutor: taskDistributor}
}
