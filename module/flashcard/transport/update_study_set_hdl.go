package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/flashcard/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateStudySetHdl godoc
// @Summary Update a study set
// @Description Update a study set
// @Tags study-set
// @Accept json
// @Produce json
// @Param StudySet body model.StudySet true "StudySet"
// @Param Authorization header string true "Bearer + Access Token"
// @Success 200 {object} model.StudySet
// @Failure 400 {object} common.ErrResp
// @Failure 403 {object} common.ErrResp
// @Failure 500 {object} common.ErrResp
// @Router /study-set [put]
func (h httpHandler) UpdateStudySetHdl() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		UserID := ctx.Value(common.CtxValUserIdKey).(uuid.UUID)

		var studySet model.StudySet

		if err := ctx.ShouldBindJSON(&studySet); err != nil {
			ctx.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, "Invalid request body", err))
			return
		}

		if studySet.UserID != UserID {
			ctx.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, "Forbidden", nil))
			return
		}

		updatedStudySet, err := h.useCase.UpdateStudySet(ctx.Request.Context(), &studySet)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, "Failed to update study set", err))
			return
		}

		ctx.JSON(http.StatusOK, updatedStudySet)

	}
}
