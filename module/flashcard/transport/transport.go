package transport

import (
	"quizen/component/token"
	"quizen/middleware"
	"quizen/module/flashcard/usecase"
	"quizen/module/user/store"

	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	useCase usecase.UseCase
}

func NewHTTPHandler(useCase usecase.UseCase) httpHandler {
	return httpHandler{useCase: useCase}
}

func InitializeFlashcardRoutes(h httpHandler, router *gin.RouterGroup, tokenprovider token.TokenProvider, store store.Store) {
	auth := middleware.Auth(tokenprovider, store)
	router.POST("/study-set", auth, h.CreateStudySetHandler())
}
