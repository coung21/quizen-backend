package main

import (
	"log"
	"quizen/component/worker"
	"quizen/config"
	"quizen/db"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func main() {
	r := gin.Default()

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
