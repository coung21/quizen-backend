package transport

import (
	"quizen/module/flashcard/usecase"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	useCase usecase.UseCase
}

func NewHTTPHandler(useCase usecase.UseCase) httpHandler {
	return httpHandler{useCase: useCase}
}

func InitializeFlashcardRoutes(h httpHandler, router *gin.RouterGroup) {
	router.POST("/study-set", h.CreateStudySetHandler())
}
