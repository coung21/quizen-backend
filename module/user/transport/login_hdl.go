package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/user/model"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	User  model.User `json:"user"`
	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"token"`
}

func (h httpHandler) LoginHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq LoginReq

		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		user, tokens, err := h.userUsecase.Login(c.Request.Context(), loginReq.Email, loginReq.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		resp := LoginResp{
			User: *user,
			Token: struct {
				AccessToken  string `json:"access_token"`
				RefreshToken string `json:"refresh_token"`
			}{
				AccessToken:  tokens.AccessToken,
				RefreshToken: tokens.RefreshToken,
			},
		}

		c.JSON(http.StatusOK, resp)
	}

}
