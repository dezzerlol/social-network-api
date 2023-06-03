package posts

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images"]
		body := form.Value["body"][0]

		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		err := h.postsService.CreatePost(ctx, files, body)

		if err != nil {
			h.payload.BadRequest(c, err)
			return
		}

		h.payload.WriteJSON(c, http.StatusCreated, "ok")
	}

}
