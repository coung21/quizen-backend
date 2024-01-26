package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

func (h httpHandler) VerifyEmailHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get email and code from query string
		email := c.Query("email")
		code := c.Query("code")

		//call usecase
		_, err := h.userUsecase.VerifyEmail(c.Request.Context(), email, code)

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
		} else {
			//response a string "success"
			c.JSON(http.StatusOK, "success")
		}
	}
}
