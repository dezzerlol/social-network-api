package auth

import (
	"context"
	"errors"
	"net/http"
	"social-network-api/internal/repository/users"
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
			switch {
			case errors.Is(err, users.ErrRecordNotFound):
				h.payload.InvalidCredentials(c)
				return
			default:
				h.payload.InternalServerError(c, err)
				return
			}
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

		token, err := user.Password.GenerateAuthToken(user.Id, 7*24*time.Hour)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		sessionUser := map[string]int64{
			"UserID": user.Id,
		}

		// save token in cache
		err = h.cache.SetStruct(ctx, token.Plaintext, sessionUser, 24*time.Hour)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		payload := map[string]interface{}{
			"auth_token": token,
		}

		c.SetCookie("auth_token", token.Plaintext, 7*24*60*60, "/", "localhost", false, true)
		h.payload.WriteJSON(c, http.StatusOK, payload)
	}

}
