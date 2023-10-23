package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RequireAuth(c *gin.Context) {
	// get auth token from header
	requestedToken := c.Request.Header.Get("Authorization")
	if requestedToken == "" {
		c.String(http.StatusUnauthorized, "Authorization required")
		c.Abort()
		return
	}

	// validate token
	accessedToken := os.Getenv("token")
	if len(requestedToken) < 7 || requestedToken[7:] != accessedToken {
		c.String(http.StatusForbidden, "Provided token is invalid")
		c.Abort()
		return
	}
	// continue
	c.Next()
}
