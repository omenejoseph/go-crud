package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	token_utils "github.com/omenejoseph/go-crud/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token_utils.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
