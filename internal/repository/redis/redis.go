package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gostack-labs/adminx/configs"
)

var _ Store = (*cache)(nil)

type Store interface {
	Set(c context.Context, key, value string, ttl time.Duration) error
	Get(c context.Context, key string) (string, error)
	TTL(c context.Context, key string) (time.Duration, error)
	Expire(c context.Context, key string, ttl time.Duration) bool
	ExpireAt(c context.Context, key string, ttl time.Time) bool
	Del(c context.Context, key string) bool
	Exists(c context.Context, key ...string) bool
	Incr(c context.Context, key string) int64
}

type cache struct {
	client *redis.Client
}

func New() (Store, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         configs.Config.Redis.Addr,
		Password:     configs.Config.Redis.Pass,
		DB:           configs.Config.Redis.Db,
		MaxRetries:   configs.Config.Redis.MaxRetries,
		PoolSize:     configs.Config.Redis.PoolSize,
		MinIdleConns: configs.Config.Redis.MinIdleConns,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &cache{client: client}, nil
}

func (c *cache) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	if err := c.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return err
	}
	return nil
}

func (c *cache) Get(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (c *cache) TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := c.client.TTL(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return ttl, nil
}

func (c *cache) Expire(ctx context.Context, key string, ttl time.Duration) bool {
	ok, _ := c.client.Expire(ctx, key, ttl).Result()
	return ok
}

func (c *cache) ExpireAt(ctx context.Context, key string, ttl time.Time) bool {
	ok, _ := c.client.ExpireAt(ctx, key, ttl).Result()
	return ok
}

func (c *cache) Exists(ctx context.Context, keys ...string) bool {
	if len(keys) == 0 {
		return true
	}
	value, _ := c.client.Exists(ctx, keys...).Result()
	return value > 0
}

func (c *cache) Del(ctx context.Context, key string) bool {
	if key == "" {
		return true
	}

	value, _ := c.client.Del(ctx, key).Result()
	return value > 0
}

func (c *cache) Incr(ctx context.Context, key string) int64 {
	value, _ := c.client.Incr(ctx, key).Result()
	return value
}