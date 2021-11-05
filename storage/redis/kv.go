package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/chyroc/go-ptr"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"time"
)

type KVDataClass struct {
	cw *CommonRedisWrapper
}

func (kv *KVDataClass) GetCommonRedis() *CommonRedisWrapper {
	return kv.GetCommonRedis()
}

func (kv *KVDataClass) GetRawRedis() *redis.Client {
	return kv.cw.RawClient
}

func (kv *KVDataClass) NewKVDataClass() *KVDataClass {
	return &KVDataClass{
		cw: NewCommonRedisWrapper(CommonRedisClient),
	}
}


func (kv *KVDataClass) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	_, err := kv.GetRawRedis().Set(ctx, key, val, expiration).Result()
	return err
}

func (kv *KVDataClass) MSet(ctx context.Context, kvs map[string]interface{}, expiration time.Duration) error {
	pipeline := kv.GetRawRedis().Pipeline()
	defer pipeline.Close()

	for k, v := range kvs {
		err := pipeline.Set(ctx, k, v, expiration).Err()
		if err != nil {
			Log.Warnf(fmt.Sprintf("Operation Set failef, err: %s", err))
		}
	}

	if _, err := pipeline.Exec(ctx); err != nil {
		Log.Errorf(fmt.Sprintf("Pipeline commit error err: %v", err))
		return errors.New(fmt.Sprintf("Pipeline commit error err: %v", err))
	}

	return nil
}

func (kv *KVDataClass) Get(ctx context.Context, key string) (*string, error) {
	v, err := kv.GetRawRedis().Get(ctx, key).Result()
	if err != nil {
		Log.Errorf("Get failed, errL %s", err)
		return nil, errors.New(fmt.Sprintf("Get failed, errL %s", err))
	}
	return ptr.String(v), nil
}

func (kv *KVDataClass) MGet(ctx context.Context, keys []string) ([][]byte, error) {
	vals, err := kv.GetRawRedis().MGet(ctx, keys...).Result()
	if err != nil {
		Log.Errorf("MGet error err: %v", err)
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

func (kv *KVDataClass) MDelete(ctx context.Context, keys []string) error {
	if _, err := kv.GetRawRedis().Del(ctx, keys...).Result(); err != nil {
		Log.Errorf("mDelete error err: %v", err)
		return errors.New(fmt.Sprintf("mDelete error err: %v", err))
	}
	return nil
}

func (kv *KVDataClass) MSetNx(ctx context.Context, kvs map[string]interface{}, expiration time.Duration) ([]string, error) {
	pipeline := kv.GetRawRedis().Pipeline()
	defer pipeline.Close()

	var failedKeys []string
	for k, v := range kvs {
		boolCmd, err := pipeline.SetNX(ctx, k, v, expiration).Result()
		if err != nil {
			Log.Warnf(fmt.Sprintf("MSetNx to redis failed, err: %s", err))
		}
		if !boolCmd {
			Log.Logf(logrus.InfoLevel, fmt.Sprintf("SetNX key: %s failed because of existence, ttl: %v", k, kv.GetCommonRedis().GetExpireTime(ctx, k)))
			failedKeys = append(failedKeys, k)
		}
	}

	_, err := pipeline.Exec(ctx)
	if err != nil {
		Log.Logf(logrus.ErrorLevel, "Commit error: %s", err)
		return nil, err
	}

	return failedKeys, nil
}
