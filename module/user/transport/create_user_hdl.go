package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/user/model"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CreateUserHandler godoc
// @Summary Register a new user
// @Description Register a new user and send a verification email
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param user body createUserRequest true "User info"
// @Success 201 {object} model.User "Success"
// @Failure 400 {object} common.ErrResp "Bad Request"
// @Failure 500 {object} common.ErrResp "Internal Server Error"
// @Router /users/register [post]
func (h *httpHandler) CreateUserHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestParams createUserRequest

		if err := c.ShouldBindJSON(&requestParams); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
			return
		}

		user := model.User{
			Username: requestParams.Username,
			Email:    requestParams.Email,
			Password: requestParams.Password,
		}

		createdUser, err := h.userUsecase.CreateUser(c.Request.Context(), &user)
		if err != nil {
			if err == common.BadRequest {
				c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else if err == common.ExistsEmailError {
				c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			}
			c.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
			return
		}

		createdUser.Sanitize()

		c.JSON(http.StatusCreated, createdUser)
	}
}
