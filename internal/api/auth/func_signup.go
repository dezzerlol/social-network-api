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

		if err := c.ShouldBindJSON(&input); err != nil {
			h.payload.BadRequest(c, err)
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
			h.payload.BadRequest(c, err)
			return
		}

		payload := map[string]interface{}{
			"message": "User created successfully.",
			"user_id": user.Id,
		}

		h.payload.WriteJSON(c, http.StatusCreated, payload)
	}

}
