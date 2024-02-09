package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

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
