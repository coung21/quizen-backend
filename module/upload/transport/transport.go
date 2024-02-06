package transport

import (
	"quizen/module/upload/usecase"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	uc usecase.Usecase
}

func NewUploadHandler(uc usecase.Usecase) httpHandler {
	return httpHandler{uc: uc}
}

func InitialzeUploadRoutes(hdl httpHandler, router *gin.RouterGroup) {
	router.POST("/upload", hdl.UploadHdl())
}
