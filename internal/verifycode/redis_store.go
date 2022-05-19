package verifycode

import (
	"context"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/repository/redis"
)

var _ Store = (*redisStore)(nil)

type redisStore struct {
	cache redis.Store
}

type Store interface {
	// save verifycode
	Set(id string, value string) error

	// get verifycode
	Get(id string, clear bool) (string, error)

	// check verifycode
	Check(id, answer string, clear bool) bool
}

func (v *redisStore) Set(key string, value string) error {

	expireTime := configs.Get().VerifyCode.ExpireTime

	return v.cache.Set(context.Background(), configs.Get().VerifyCode.KeyPrefix+key, value, expireTime)
}

func (v *redisStore) Get(key string, clear bool) (string, error) {
	key = configs.Get().VerifyCode.KeyPrefix + key
	val, err := v.cache.Get(context.Background(), key)
	if err != nil {
		//log.Fatal("get verifycode err:", err)
		return "", err
	}
	if clear {
		_ = v.cache.Del(context.Background(), key)
	}
	return val, nil
}

func (v *redisStore) Check(key, answer string, clear bool) bool {
	val, err := v.Get(key, clear)
	if err != nil {
		return false
	}
	return val == answer
}
