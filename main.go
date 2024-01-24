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

	userStore := userstore.NewUserStore(mdb)

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	_ = worker.NewRedisTaskDistributor(asynq.RedisClientOpt{Addr: config.RedisAddress})
	go worker.RunTaskProcessor(asynq.RedisClientOpt{Addr: config.RedisAddress}, userStore, mailer)

	r.Run(config.ServerAddress)
}
