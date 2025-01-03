package middlewares

import (
	"net/http"
	"strings"
	"todo-app/internal/constants"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userService services.IUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := userService.GetUserFromToken(tokenString)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}

		c.Set(constants.CTX_USER, user)

		c.Next()
	}
}
