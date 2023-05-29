package posts

import "github.com/gin-gonic/gin"

func (h *handler) RemoveComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post",
		})
	}
}
