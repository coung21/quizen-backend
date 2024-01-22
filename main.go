package main

import (
	"quizen/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Server struct {
	Engine *gin.Engine
}

func (s *Server) Init() {
	s.Engine = gin.Default()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	_, err := db.Connect()
	if err != nil {
		panic(err)
	}

}

func main() {
	s := Server{}
	s.Init()

	s.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	s.Engine.Run(":8080")
}
