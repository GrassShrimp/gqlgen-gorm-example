package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const apqPrefix = "apq:"

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

func NewCache(redisAddress string, password string, db int, ttl time.Duration) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: password,
		DB:       db,
	})

	err := client.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("could not create cache: %w", err)
	}

	return &Cache{client: client, ttl: ttl}, nil
}

func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(apqPrefix+key, value, c.ttl)
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(apqPrefix + key).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
