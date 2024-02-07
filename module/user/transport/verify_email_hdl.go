package transport

import (
	"net/http"
	"quizen/common"

	"github.com/gin-gonic/gin"
)

// VerifyEmailHdl godoc
// @Summary Verify email
// @Description Verify email with code
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param email query string true "Email"
// @Param code query string true "Code"
// @Success 200 {string} string "Success"
// @Failure 500 {object} common.ErrResp "Internal Server Error"
// @Router /users/verify-email [get]
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
