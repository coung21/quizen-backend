package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/user/model"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestResp(http.StatusBadRequest, err.Error(), err))
			return
		}

		createdUser, err := h.userUsecase.CreateUser(c.Request.Context(), user)
		if err != nil {
			if err == common.BadRequest {
				c.JSON(http.StatusBadRequest, common.NewRestResp(http.StatusBadRequest, err.Error(), err))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewRestResp(http.StatusInternalServerError, err.Error(), err))
			return
		}

		createdUser.Sanitize()

		c.JSON(http.StatusOK, common.NewRestResp(http.StatusOK, "success", createdUser))
	}
}
