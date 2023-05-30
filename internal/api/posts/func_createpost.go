package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images"]
		body := form.Value["body"][0]

		err := h.postsService.CreatePost(files, body)

		if err != nil {
			h.payload.BadRequest(c, err)
			return
		}

		h.payload.WriteJSON(c, http.StatusCreated, "ok")
	}
}
