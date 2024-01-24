package main

import (
	"log"
	"os"
	"quizen/component/mail"
	"quizen/component/worker"
	"quizen/config"
	"quizen/db"
	"quizen/middleware"
	userstore "quizen/module/user/store"
	userTransport "quizen/module/user/transport"
	useruc "quizen/module/user/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zlog.Logger = zlog.Output(output)

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mdb, err := db.Connect(
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlDb,
		config.MysqlHost,
		config.MysqlPort,
	)

	if err != nil {
		panic(err)
	}

	taskDistributor := worker.NewRedisTaskDistributor(asynq.RedisClientOpt{Addr: config.RedisAddress})

	userStore := userstore.NewUserStore(mdb)
	userUc := useruc.NewUserUsecase(userStore, taskDistributor)
	userTransport.InitializeUserRoutes(userTransport.NewHTTPHandler(userUc), r.Group("/v1/users"))

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	go worker.RunTaskProcessor(asynq.RedisClientOpt{Addr: config.RedisAddress}, userStore, mailer)

	r.Run(config.ServerAddress)
}
