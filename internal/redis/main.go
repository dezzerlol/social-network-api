package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func New(host string, port int, pass string) *Client {
	return &Client{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: pass,
			DB:       0,
		}),
	}
}
