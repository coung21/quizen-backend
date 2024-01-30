package transport

import (
	"context"
	"net/http"
	"quizen/common"
	"quizen/module/user/model"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	User      model.User `json:"user"`
	SessionID string     `json:"session_id"`
	Token     struct {
		AccessToken     string    `json:"access_token"`
		RefreshToken    string    `json:"refresh_token"`
		RefreshTokenExp time.Time `json:"refresh_token_exp"`
	} `json:"token"`
}

// LoginHdl godoc
// @Summary Login
// @Description Login
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param user body LoginReq true "email and password"
// @Success 200 {object} LoginResp "Success"
// @Failure 400 {object} common.ErrResp "Bad Request"
// @Failure 500 {object} common.ErrResp "Internal Server Error"
// @Router /users/login [post]
func (h httpHandler) LoginHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq LoginReq

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		ctx := c.Request.Context()

		ctx = context.WithValue(ctx, "user_agent", c.Request.UserAgent())
		ctx = context.WithValue(ctx, "user_ip", c.ClientIP())

		user, tokens, sessionID, err := h.userUsecase.Login(ctx, loginReq.Email, loginReq.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		user.Sanitize()

		resp := LoginResp{
			User:      *user,
			SessionID: sessionID,
			Token: struct {
				AccessToken     string    `json:"access_token"`
				RefreshToken    string    `json:"refresh_token"`
				RefreshTokenExp time.Time `json:"refresh_token_exp"`
			}{
				AccessToken:     tokens.AccessToken,
				RefreshToken:    tokens.RefreshToken,
				RefreshTokenExp: tokens.RefreshTokenExp,
			},
		}

		c.JSON(http.StatusOK, resp)
	}

}
