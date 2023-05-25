package http

import (
	"social-network-api/internal/api/auth"

	"github.com/gin-gonic/gin"
)

func (s *Server) setHTTPRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	authHandler := auth.New(s.logger, s.db, s.cache)

	v1 := router.Group("/v1")
	{
		// AUTH
		auth := v1.Group("/auth")
		auth.POST("/signup", authHandler.Signup())
		auth.POST("/login", authHandler.Login())
	}

	return router
}
