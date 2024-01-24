package usecase

import (
	"context"
	"quizen/common"
	"quizen/component/worker"
	"quizen/module/user/model"
	"time"

	"github.com/hibiken/asynq"
)

func (uc userUsecase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	fUser, _ := uc.userStore.GetUserByEmail(ctx, user.Email)
	if fUser != nil {
		return nil, common.ExistsEmailError
	}

	user.HashPassword()

	nUser, err := uc.userStore.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	taskPayload := &worker.PayloadVerifyEmail{
		Username: user.Username,
		Email:    user.Email,
		Code:     common.RandomString(10),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}

	defer uc.taskDistrbutor.DistrbuteTaskSendVerifyEmail(ctx, taskPayload, opts...)

	return nUser, nil
}
