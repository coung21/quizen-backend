package mock

import (
	"context"
	"quizen/component/worker"

	"github.com/hibiken/asynq"
)

type TaskDistributorMock struct {
	DistributeTaskSendVerifyEmailFn func(ctx context.Context, payload *worker.PayloadVerifyEmail, opts ...asynq.Option) error
}

func (m *TaskDistributorMock) DistrbuteTaskSendVerifyEmail(
	ctx context.Context,
	payload *worker.PayloadVerifyEmail,
	opts ...asynq.Option,
) error {
	return m.DistributeTaskSendVerifyEmailFn(ctx, payload, opts...)
}
