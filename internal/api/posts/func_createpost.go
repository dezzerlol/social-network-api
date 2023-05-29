package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["images"]
		//body := form.Value["body"][0]

		for _, file := range files {
			c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		}

		h.payload.WriteJSON(c, http.StatusCreated, "ok")
	}
}
