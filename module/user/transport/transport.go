package transport

import (
	"quizen/module/user/usecase"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	userUsecase usecase.Usecase
}

func NewHTTPHandler(userUsecase usecase.Usecase) httpHandler {
	return httpHandler{userUsecase: userUsecase}
}

func InitializeUserRoutes(hdl httpHandler, router *gin.RouterGroup) {
	router.POST("/register", hdl.CreateUserHandler())
	router.GET("/verify-email", hdl.VerifyEmailHdl())
	router.POST("/login", hdl.LoginHdl())
}
