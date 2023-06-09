package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		strid := c.Query("id")
		id, err := strconv.Atoi(strid)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		err = h.postsService.DeletePost(ctx, id)

		if err != nil {
			h.payload.InternalServerError(c, err)
			return
		}

		h.payload.WriteJSON(c, 200, "ok")
	}
}
