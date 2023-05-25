package auth

import (
	"context"
	"net/http"
	"social-network-api/internal/db/models"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email     string `json:"email" binding:"required"`
			Password  string `json:"password" binding:"required"`
			Username  string `json:"username" binding:"required"`
			Firstname string `json:"firstname" binding:"required"`
			Lastname  string `json:"lastname" binding:"required"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		user := &models.User{
			Email:     input.Email,
			Username:  input.Username,
			Firstname: input.Firstname,
			Lastname:  input.Lastname,
		}
		user.Password.PlainTextPass = input.Password

		if err := h.userService.Create(ctx, user); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully.",
			"user_id": user.Id,
		})
	}

}
