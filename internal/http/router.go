package http

import (
	"social-network-api/internal/api/auth"
	"social-network-api/internal/api/posts"

	"github.com/gin-gonic/gin"
)

func (s *Server) setHTTPRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.MaxMultipartMemory = 32 << 20 // 32 mb

	authHandler := auth.New(s.logger, s.db, s.cache)
	postsHandler := posts.New(s.logger, s.db, s.cache)

	v1 := router.Group("/v1")
	{
		// AUTH
		auth := v1.Group("/auth")
		auth.POST("/signup", authHandler.Signup())
		auth.POST("/login", authHandler.Login())

		auth.POST("/logout", authHandler.Logout())
	}

	{
		// POSTS
		posts := v1.Group("/posts")
		posts.Use(s.AuthSession())

		posts.POST("/", postsHandler.CreatePost())
		posts.DELETE("/:id", postsHandler.DeletePost())
		posts.POST("/:id/like", postsHandler.Like())
		posts.DELETE("/:id/like", postsHandler.RemoveLike())
		posts.POST("/:id/comment", postsHandler.Comment())
		posts.DELETE("/:id/comment/:comment_id", postsHandler.RemoveComment())
	}

	return router
}
