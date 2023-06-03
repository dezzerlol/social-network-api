package redis

import (
	"encoding/json"
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

func (c Client) MarshalBinary(val interface{}) (data []byte, err error) {
	bytes, err := json.Marshal(val)
	return bytes, err
}
