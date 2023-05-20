package models

import (
	"time"
)

type Post struct {
	Id         string      `json:"id"`
	UserId     string      `json:"user_id"`
	Body       string      `json:"body"`
	Images     *ImageJsonb `json:"images"`
	Created_at time.Time   `json:"created_at"`
}
