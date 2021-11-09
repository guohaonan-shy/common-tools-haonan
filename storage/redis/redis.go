package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type ListOperationDirection int64

const (
	ListPushFromLeft     ListOperationDirection = 1
	ListPushFromRight    ListOperationDirection = 2
	ListInsertFromBefore ListOperationDirection = 3
	ListInsertFromAfter  ListOperationDirection = 4
)

var (
	CommonRedisClient *redis.Client

	Log = logrus.New()
)

func CommonRedisWrapperInit(address string) (*CommonRedisWrapper, error) {
	CommonRedisClient = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err := CommonRedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return NewCommonRedis(CommonRedisClient), nil
}

type CommonRedisWrapper struct {
	RawClient *redis.Client
}

func NewCommonRedis(client *redis.Client) *CommonRedisWrapper {
	return &CommonRedisWrapper{
		RawClient: client,
	}
}

type ConnectAcquire interface {
	GetCommonRedis() *CommonRedisWrapper
	GetRawRedis() *redis.Client
}


func (c *CommonRedisWrapper) Lock(ctx context.Context, key string, uuid string, timeout, retry time.Duration) error {
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

func (c *CommonRedisWrapper) SetExpireTime(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return c.RawClient.Expire(ctx, key, expiration).Result()
}

func (c *CommonRedisWrapper) GetExpireTime(ctx context.Context, key string) time.Duration {
	return c.RawClient.TTL(ctx, key).Val()
}