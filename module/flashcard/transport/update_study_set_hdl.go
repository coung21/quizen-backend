package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/flashcard/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h httpHandler) UpdateStudySetHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		UserID := ctx.Value(common.CtxValUserIdKey).(uuid.UUID)

		var studySet model.StudySet

		if err := ctx.ShouldBindJSON(&studySet); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestResp(http.StatusBadRequest, "Invalid request body", err))
			return
		}

		if studySet.UserID != UserID {
			ctx.JSON(http.StatusForbidden, common.NewRestResp(http.StatusForbidden, "Forbidden", nil))
			return
		}

		updatedStudySet, err := h.useCase.UpdateStudySet(ctx.Request.Context(), &studySet)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewRestResp(http.StatusInternalServerError, "Failed to update study set", err))
			return
		}

		ctx.JSON(http.StatusOK, updatedStudySet)

	}
}
