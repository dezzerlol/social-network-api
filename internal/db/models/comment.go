package models

import "time"

type Comment struct {
	PostId     string      `json:"post_id"`
	UserId     string      `json:"user_id"`
	Body       string      `json:"body"`
	Images     *ImageJsonb `json:"images"`
	Created_at time.Time   `json:"created_at"`
}
