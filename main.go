package main

import (
	"log"
	"os"
	"quizen/component/worker"
	"quizen/config"
	"quizen/db"
	"quizen/middleware"
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

	_, err = db.Connect(config.MySqlUri)

	if err != nil {
		panic(err)
	}

	_ = worker.NewRedisTaskDistributor(asynq.RedisClientOpt{Addr: config.RedisAddress})
	go worker.RunTaskProcessor(asynq.RedisClientOpt{Addr: config.RedisAddress})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.ServerAddress)
}
