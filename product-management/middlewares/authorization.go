package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorize *****, request type: ", c.Request.Method)

		role := c.GetString("user_type")
		fmt.Println("Authorization, user roles: ", role)

		if !strings.EqualFold(c.Request.Method, "GET") && !strings.EqualFold(role, "admin") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("User is not authorized")})
			c.Abort()
			return
		}

		c.Next()
	}
}
