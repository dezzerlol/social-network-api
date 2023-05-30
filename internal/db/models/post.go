package models

import (
	"time"
)

type Post struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	Body       string    `json:"body"`
	Images     []Media   `json:"media"`
	Created_at time.Time `json:"created_at"`
}
