package main

import (
	"log"
	"quizen/config"
	"quizen/db"

	"github.com/gin-gonic/gin"
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(config.ServerAddress)
}
