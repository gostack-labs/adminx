package redis

import (
	"context"
	"strings"
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
	Version() string
}

type cache struct {
	client *redis.Client
}

func New() (Store, error) {
	var redisConf = configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Pass,
		DB:           redisConf.Db,
		MaxRetries:   redisConf.MaxRetries,
		PoolSize:     redisConf.PoolSize,
		MinIdleConns: redisConf.MinIdleConns,
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

func (c *cache) Version() string {
	server := c.client.Info(context.Background(), "server").Val()
	spl1 := strings.Split(server, "# Server")
	spl2 := strings.Split(spl1[1], "redis_version:")
	spl3 := strings.Split(spl2[1], "redis_git_sha1:")
	return spl3[0]
}
