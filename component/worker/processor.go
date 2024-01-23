package worker

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/hibiken/asynq"
)

type Processor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, t *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt) Processor {
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
	return &RedisTaskProcessor{server: server}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, p.ProcessTaskSendVerifyEmail)

	return p.server.Run(mux)
}

func RunTaskProcessor(redisOpt asynq.RedisClientOpt) {
	processor := NewRedisTaskProcessor(redisOpt)
	log.Info().Time("time", time.Now()).Msg("task processor starting")
	err := processor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
