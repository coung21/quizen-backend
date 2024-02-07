package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

type renewTokenReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
	SessionID    string `json:"session_id" binding:"required"`
}

type renewTokenResp struct {
	AccessToken string `json:"access_token"`
}

// RenewTokenHdl godoc
// @Summary Renew access token
// @Description Renew access token
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param ususer body renewTokenReq true "User info"
// @Success 200 {object} renewTokenResp "Success"
// @Failure 400 {object} common.ErrResp "Bad Request"
// @Failure 401 {object} common.ErrResp "Unauthorized"
// @Failure 403 {object} common.ErrResp "Forbidden"
// @Failure 404 {object} common.ErrResp "Not Found"
// @Failure 500 {object} common.ErrResp "Internal Server Error"
// @Router /users/renew-token [post]
func (h httpHandler) RenewTokenHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req renewTokenReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, "invalid request body", nil))
			return
		}

		accessToken, err := h.userUsecase.RenewToken(c.Request.Context(), req.SessionID, req.RefreshToken)

		if err != nil {
			if err == common.NotFound {
				c.JSON(http.StatusNotFound, common.NewRestErr(http.StatusNotFound, "session not found", nil))
				return
			} else if err == common.InvalidJWTToken {
				c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, "invalid refresh token", nil))
				return
			} else if err == common.Forbidden {
				c.JSON(http.StatusForbidden, common.NewRestErr(http.StatusForbidden, "invalid refresh token", nil))
				return
			} else if err == common.ErrJWTExpired {
				c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, "refresh token expired", nil))
				return
			} else if err == common.Unauthorized {
				c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, "invalid refresh token", nil))
				return
			} else {
				c.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, "unexpected error", nil))
				return
			}
		}

		c.JSON(http.StatusOK, renewTokenResp{AccessToken: accessToken})
	}
}
