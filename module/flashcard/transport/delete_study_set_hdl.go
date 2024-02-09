package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

// DeleteStudySetHdl godoc
// @Summary Delete a study set
// @Description Delete a study set
// @Tags study-set
// @Produce json
// @Param id path string true "Study set ID"
// @Param Authorization header string true "Bearer + Access Token"
// @Success 204 {string} string "No content"
// @Failure 404 {object} common.ErrResp "Study set not found"
// @Failure 500 {object} common.ErrResp "Failed to delete study set"
// @Router /study-set/{id} [delete]
func (h httpHandler) DeleteStudySetHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		studySetID := c.Param("id")
		if err := h.useCase.DeleteStudySet(c.Request.Context(), studySetID); err != nil {
			if err == common.NotFound {
				c.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, "Study set not found", err))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, "Failed to delete study set", err))
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}
