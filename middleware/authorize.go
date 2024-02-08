package middleware

import (
	"net/http"
	"quizen/common"
	"quizen/component/token"
	"quizen/module/user/store"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeader(s string) (string, error) {
	parts := strings.Split(s, " ")

	if len(parts) != 2 || parts[0] != "Bearer" || parts[1] == "" {
		return "", common.InvalidAuthHeader
	}

	return parts[1], nil
}

func Auth(tokenprovider token.TokenProvider, store store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, err.Error(), nil))
			c.Abort()
			return
		}

		payload, err := tokenprovider.Validate(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.InvalidJWTToken.Error(), nil))
			c.Abort()
			return
		}

		user, err := store.GetUser(c.Request.Context(), map[string]interface{}{"id": payload.ID.String()})

		if err != nil {
			c.JSON(http.StatusUnauthorized, common.NewRestErr(http.StatusUnauthorized, common.InvalidJWTToken.Error(), nil))
			c.Abort()
			return
		}

		c.Set(common.CtxValUserIdKey, user.ID)
		c.Next()
	}
}
