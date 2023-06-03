package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) AuthSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Cookie("auth_token")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "session cookie not found"})
			return
		}

		sessionId := s.cache.Get(c.Request.Context(), authCookie)

		if sessionId == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Set("auth_token", sessionId)
		c.Next()
	}
}
