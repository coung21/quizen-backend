package transport

import (
	"quizen/component/token"
	"quizen/module/user/store"
	"quizen/module/user/usecase"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	userUsecase usecase.Usecase
}

func NewHTTPHandler(userUsecase usecase.Usecase) httpHandler {
	return httpHandler{userUsecase: userUsecase}
}

func InitializeUserRoutes(hdl httpHandler, router *gin.RouterGroup, tokenprovider token.TokenProvider, store store.Store) {
	router.POST("/register", hdl.CreateUserHandler())
	router.GET("/verify-email", hdl.VerifyEmailHdl())
	router.POST("/login", hdl.LoginHdl())
	router.DELETE("/logout/:session_id", hdl.LogoutHdl())
	router.POST("/renew-token", hdl.RenewTokenHdl())
}
