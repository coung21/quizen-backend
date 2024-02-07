package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

// LogoutHdl godoc
// @Summary Logout
// @Description Logout
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param session_id path string true "session_id"
// @Success 200 {object} string "Success"
// @Failure 400 {object} common.ErrResp "Bad Request"
// @Failure 404 {object} common.ErrResp "Not Found"
// @Failure 500 {object} common.ErrResp "Internal Server Error"
// @Router /users/logout/{session_id} [delete]
func (h httpHandler) LogoutHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := string(c.Param("session_id"))

		if sessionID == "" {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, "session_id is required", nil))
			return
		}

		if err := h.userUsecase.Logout(c.Request.Context(), sessionID); err != nil {
			if err == common.NotFound {
				c.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, "session not found", nil))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}

		c.JSON(http.StatusOK, common.NewRestResp(http.StatusOK, "Logout success", nil))
	}
}
