package payload

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Payload struct {
	logger *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger) *Payload {
	return &Payload{
		logger: logger,
	}
}

func (p *Payload) WriteJSON(c *gin.Context, status int, payload interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, map[string]interface{}{
		"data": payload,
	})
}

func (p *Payload) BadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

func (p *Payload) InternalServerError(c *gin.Context, err error) {
	msg := "The server encountered a problem and could not process your request"

	p.logger.Errorln(err, map[string]interface{}{
		"req_method": c.Request.Method,
		"req_url":    c.Request.URL,
	})

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": msg,
	})
}

func (p *Payload) NotFound(c *gin.Context) {
	msg := "The requested resource could not be found"

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error": msg,
	})
}

func (p *Payload) Unauthorized(c *gin.Context) {
	msg := "You are not authorized to access this resource"

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
}

func (p *Payload) InvalidCredentials(c *gin.Context) {
	msg := "Invalid credentials"

	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": msg,
	})
}
