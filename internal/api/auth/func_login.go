package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		user, err := h.userService.FindByEmail(ctx, input.Email)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		passMatch, err := user.Password.Matches(input.Password)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		if !passMatch {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		token, err := user.Password.GenerateAuthToken(user.Id, 24*time.Hour)

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// save token in cache
		err = h.cache.Set(ctx, token.Plaintext, user.Id, 24*time.Hour).Err()

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"auth_token": token,
		})
	}

}
