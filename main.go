package main

import (
	"log"
	"os"
	"quizen/component/mail"
	"quizen/component/token"
	"quizen/component/worker"
	"quizen/config"
	"quizen/db"
	"quizen/middleware"
	userstore "quizen/module/user/store"
	userTransport "quizen/module/user/transport"
	useruc "quizen/module/user/usecase"
	"time"

	_ "quizen/docs"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Quizen API
// @description This is a flashcard learning app API.
// @version 1.0
// @Host localhost:8080
// @BasePath /v1

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zlog.Logger = zlog.Output(output)

	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())

	//Add swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	tokenProvider := token.NewJWTProvider(config.SecretKey, config.AccessTokenDuration, config.RefreshTokenDuration)
	taskDistributor := worker.NewRedisTaskDistributor(asynq.RedisClientOpt{Addr: config.RedisAddress})

	userStore := userstore.NewUserStore(mdb)
	userUc := useruc.NewUserUsecase(userStore, taskDistributor, tokenProvider)
	userTransport.InitializeUserRoutes(userTransport.NewHTTPHandler(userUc), r.Group("/v1/users"))

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	go worker.RunTaskProcessor(asynq.RedisClientOpt{Addr: config.RedisAddress}, userStore, mailer)

	r.Run(config.ServerAddress)
}
