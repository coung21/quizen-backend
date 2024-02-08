package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"quizen/component/cloudstorage"
	"quizen/component/mail"
	"quizen/component/token"
	"quizen/component/worker"
	"quizen/config"
	"quizen/db"
	"quizen/middleware"
	flashcardstore "quizen/module/flashcard/store"
	flashcardTransport "quizen/module/flashcard/transport"
	flashcardUsecase "quizen/module/flashcard/usecase"
	uploadTansport "quizen/module/upload/transport"
	uploadUsecase "quizen/module/upload/usecase"
	userstore "quizen/module/user/store"
	userTransport "quizen/module/user/transport"
	useruc "quizen/module/user/usecase"
	"syscall"
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

	srv := http.Server{
		Addr:    config.ServerAddress,
		Handler: r,
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

	s3Provider := cloudstorage.NewS3Storage(config.S3BucketName, config.S3Region, config.S3AccessKey, config.S3SecretKey, config.S3Domain)
	uploadUc := uploadUsecase.NewUploadUc(s3Provider)
	uploadTansport.InitialzeUploadRoutes(uploadTansport.NewUploadHandler(uploadUc), r.Group("/v1/uploads"))

	flashcardStore := flashcardstore.NewFlashcardStore(mdb)
	flashcardUc := flashcardUsecase.NewFlashcardUseCase(flashcardStore)
	flashcardTransport.InitializeFlashcardRoutes(flashcardTransport.NewHTTPHandler(flashcardUc), r.Group("/v1/"))

	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	go worker.RunTaskProcessor(asynq.RedisClientOpt{Addr: config.RedisAddress}, userStore, mailer)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}
}
