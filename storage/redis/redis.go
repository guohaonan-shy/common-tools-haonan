package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	CommonRedisClient *redis.Client
)

func CommonRedisWrapperInit() *CommonRedisWrapper {
	CommonRedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return NewCommonRedisWrapper(CommonRedisClient)
}

type CommonRedisWrapper struct {
	RawClient *redis.Client
	Logger    *logrus.Logger
}

func NewCommonRedisWrapper(client *redis.Client) *CommonRedisWrapper {
	return &CommonRedisWrapper{
		RawClient: client,
		Logger:    logrus.New(),
	}
}

func (c *CommonRedisWrapper) MSet(ctx context.Context, kvs map[string]interface{}, expiration time.Duration) error {
	pipeline := c.RawClient.Pipeline()
	defer pipeline.Close()

	for k, v := range kvs {
		if _, err := pipeline.Set(ctx, k, v, expiration).Result(); err != nil {
			c.Logger.Logf(logrus.ErrorLevel, fmt.Sprintf("Set error err: %v", err))
			return errors.New(fmt.Sprintf("Set error err: %v", err))
		}
	}

	if _, err := pipeline.Exec(ctx); err != nil {
		c.Logger.Logf(logrus.ErrorLevel, fmt.Sprintf("Pipeline commit error err: %v", err))
		return errors.New(fmt.Sprintf("Pipeline commit error err: %v", err))
	}

	return nil
}


func (c *CommonRedisWrapper) MGet(ctx context.Context, keys []string) ([][]byte, error) {
	vals, err := c.RawClient.MGet(ctx, keys...).Result()
	if err != nil {
		c.Logger.Logf(logrus.WarnLevel, "MGet error err: %v", err)
		return nil, errors.New(fmt.Sprintf("MGet error err: %v", err))
	}

	ret := make([][]byte, 0, len(vals))
	for _, val := range vals {
		switch v := val.(type) {
		case []byte:
			ret = append(ret, v)
		case string:
			ret = append(ret, []byte(v))
		case nil:
			ret = append(ret, nil)
		}
	}

	return ret, nil
}

func (c *CommonRedisWrapper) MDelete(ctx context.Context, keys []string) error {
	if _, err := c.RawClient.Del(ctx, keys...).Result(); err != nil {
		c.Logger.Logf(logrus.WarnLevel, "mDelete error err: %v", err)
		return errors.New(fmt.Sprintf("mDelete error err: %v", err))
	}
	return nil
}

func (c *CommonRedisWrapper) Lock (ctx context.Context, key string, uuid string, timeout, retry time.Duration) error{
	var (
		err    error
		isLock bool
	)
	until := time.Now().Add(retry)
	for {
		if isLock, err = c.RawClient.SetNX(ctx, key, uuid, timeout).Result(); err != nil {
			return errors.New(fmt.Sprintf("Lock failed err: %s", err))
		}
		if isLock {
			return nil
		} else if time.Now().After(until) {
			return errors.New(fmt.Sprintf("RedisLock retry %v s timeout", retry/time.Second))
		} else {
			time.Sleep(2 * time.Millisecond)
		}
	}
}

func (c *CommonRedisWrapper) Unlock(ctx context.Context, key string, uuid string) (bool, error) {
	// lua
	script := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	if result, err := c.RawClient.Eval(ctx, script, []string{key}, uuid).Result(); err != nil {
		return false, errors.New(fmt.Sprintf("RedisUnLock failed, err: %s", err))
	} else if result.(int64) != 1 {
		return false, errors.New(fmt.Sprintf("RedisUnLock failed result: %v", result))
	}
	return true, nil
}