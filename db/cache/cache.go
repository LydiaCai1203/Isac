package cache

import (
	"context"

	"framework/config"

	"github.com/redis/go-redis/v9"
)

// Cache 缓存
type Cache struct {
	*redis.Client
}

// NewCache 初始化 Cache
func NewCache(config *config.Config) (*Cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.addr"),
		Password: config.GetString("redis.password"),
	})
	// test connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &Cache{
		Client: rdb,
	}, nil
}
