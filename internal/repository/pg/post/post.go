package post

import (
	"social-network-api/internal/repository/pg/image"
	"time"
)

type Post struct {
	Id         string `json:"id"`
	Body       string
	Images     *image.ImageJsonb
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
