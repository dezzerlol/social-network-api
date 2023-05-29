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
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required,min=6"`
		}

		if err := h.payload.ReadJSON(c, &input); err != nil {
			h.payload.ValidationError(c, err)
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		user, err := h.userService.FindByEmail(ctx, input.Email)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		passMatch, err := user.Password.Matches(input.Password)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		if !passMatch {
			h.payload.InvalidCredentials(c)
			return
		}

		token, err := user.Password.GenerateAuthToken(user.Id, 24*time.Hour)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		// save token in cache
		err = h.cache.Set(ctx, token.Plaintext, user.Id, 24*time.Hour).Err()

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		payload := map[string]interface{}{
			"auth_token": token,
		}

		h.payload.WriteJSON(c, http.StatusOK, payload)
	}

}
