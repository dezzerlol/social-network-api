package context

import "github.com/gin-gonic/gin"

func SetUser(c *gin.Context, userId int64) {
	c.Set("UserID", userId)
}

func GetUser(c *gin.Context) int64 {
	userId := c.GetInt64("UserID")
	return userId
}
