package worker

import (
	"context"
	"quizen/component/mail"
	userstore "quizen/module/user/store"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
)

const (
	QueueDefault  = "default"
	QueueCritical = "critical"
)

type Processor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, t *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	ustore userstore.Store
	mailer mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store userstore.Store, mailer mail.EmailSender) Processor {
	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			"default":  5,
			"critical": 10,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Type("type", task).Bytes("payload", task.Payload()).Msg("error when processing task")
		}),
		Logger: NewLogger(),
	})
	return &RedisTaskProcessor{server: server, ustore: store, mailer: mailer}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, p.ProcessTaskSendVerifyEmail)

	return p.server.Run(mux)
}

func RunTaskProcessor(redisOpt asynq.RedisClientOpt, store userstore.Store, mailer mail.EmailSender) {
	processor := NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Time("time", time.Now()).Msg("task processor starting")
	err := processor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
