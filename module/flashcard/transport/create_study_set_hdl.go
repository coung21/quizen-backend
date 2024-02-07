package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/flashcard/model"

	"github.com/gin-gonic/gin"
)

// CreateStudySetHandler godoc
// @Summary Create a new study set
// @Description Create a new study set
// @Tags study-set
// @Accept  json
// @Produce  json
// @Param studySet body model.StudySet true "StudySet object that needs to be created"
// @Success 201 {object} model.StudySet
// @Failure 400 {object} common.ErrResp
// @Router /study-set [post]
func (h httpHandler) CreateStudySetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var studySet model.StudySet
		if err := c.ShouldBindJSON(&studySet); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, "invalid json body", nil))
			return
		}
		createdStudySet, err := h.useCase.CreateStudySet(c, &studySet)
		if err != nil {
			if err == model.ErrFlashCardLen {
				c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), nil))
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, createdStudySet)
	}
}
