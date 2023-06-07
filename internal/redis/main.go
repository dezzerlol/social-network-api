package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RClient struct {
	*redis.Client
}

func New(host string, port int, pass string) *RClient {
	return &RClient{
		redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: pass,
			DB:       0,
		}),
	}
}

func NewWithAddr(addr string, pass string) *RClient {
	return &RClient{
		redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       0,
		}),
	}
}

func (c RClient) MarshalBinary(val interface{}) (data []byte, err error) {
	bytes, err := json.Marshal(val)
	return bytes, err
}

func (c RClient) UnmarshalBinary(data []byte, val interface{}) error {
	err := json.Unmarshal(data, &val)
	return err
}

func (c RClient) SetStruct(ctx context.Context, key string, val interface{}, time time.Duration) error {
	bytes, err := c.MarshalBinary(val)
	if err != nil {
		return err
	}

	return c.Client.Set(ctx, key, bytes, time).Err()
}

func (c RClient) GetStruct(ctx context.Context, key string, val interface{}) error {
	data, err := c.Get(ctx, key).Result()

	if err != nil {
		return err
	}

	err = c.UnmarshalBinary([]byte(data), val)

	return err
}
