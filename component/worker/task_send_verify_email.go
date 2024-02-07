package worker

import (
	"context"
	"fmt"
	usermodel "quizen/module/user/model"
	"time"

	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TaskSendVerifyEmail = "task:send_verify_email"
)

type PayloadVerifyEmail struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

func (d *RedisTaskDistributor) DistrbuteTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}
	task := asynq.NewTask(
		TaskSendVerifyEmail,
		jsonPayload,
		opts...,
	)
	info, err := d.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).Str("task", info.Queue).
		Int("max_retry", info.MaxRetry).Dur("timeout", info.Timeout).Msg("enqueued task")

	return nil
}

func (p *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, t *asynq.Task) error {
	var payload PayloadVerifyEmail
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	//Todo: send email
	vEmail, err := p.ustore.CreateVerifyEmail(ctx, &usermodel.VerifyEmail{
		Email:      payload.Email,
		SecretCode: payload.Code,
		ExpriedAt:  time.Now().Add(10 * time.Minute),
	})

	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	subject := "Welcome to Quizen"
	verifyUrl := fmt.Sprintf("%s/verify-email?email=%s&code=%s", "http://localhost:8080/v1/users", vEmail.Email, vEmail.SecretCode)
	content := fmt.Sprintf("Hi %s, <br> Welcome to Quizen. <br> Please verify your email by clicking this link: <a href=\"%s\">Verify</a>", payload.Username, verifyUrl)
	to := []string{payload.Email}

	if err := p.mailer.SendEmail(subject, content, to, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Info().Str("type", t.Type()).Str("email", payload.Email).Msg("processed task")

	return nil
}
