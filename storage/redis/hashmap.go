package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type HashmapDataClass struct {
	cw *CommonRedisWrapper
}

func (h *HashmapDataClass) GetCommonRedis() *CommonRedisWrapper {
	return h.GetCommonRedis()
}

func (h *HashmapDataClass) GetRawRedis() *redis.Client {
	return h.cw.RawClient
}

func (h *HashmapDataClass) NewHashmapDataClass() *HashmapDataClass {
	return &HashmapDataClass{
		cw: NewCommonRedis(CommonRedisClient),
	}
}

// Set one field in a hashmap, such as username in user's info
// HSet("name", map[string]interface{}{"key1": "value1", "key2": "value2"})
func (h *HashmapDataClass) Set(ctx context.Context, key string, kvs map[string]interface{}) (int64, error) {
	return h.GetRawRedis().HSet(ctx, key, kvs).Result()
}

func (h *HashmapDataClass) Get(ctx context.Context, key string, field string) (string, error) {
	return h.GetRawRedis().HGet(ctx, key, field).Result()
}

func (h *HashmapDataClass) Del(ctx context.Context, key string, fields []string) (val int64, err error) {
	return h.GetRawRedis().HDel(ctx, key, fields...).Result()
}

func (h *HashmapDataClass) HIncrBy(ctx context.Context, key string, field string, increment int64) (result int64, err error) {
	return h.GetRawRedis().HIncrBy(ctx, key, field, increment).Result()
}

// GetAll get all fields under a specific key
func (h *HashmapDataClass) GetAll(ctx context.Context, key string) (map[string]string, error) {
	return h.GetRawRedis().HGetAll(ctx, key).Result()
}


func (h *HashmapDataClass) MGet(ctx context.Context, key string, fields []string) (map[string]interface{}, error) {
	vals, err := h.GetRawRedis().HMGet(ctx, key, fields...).Result()
	if err != nil {
		Log.Errorf(fmt.Sprintf("HashMap MGet failed, err: %s", err))
		return nil, err
	}

	kvs := make(map[string]interface{}, len(fields))
	for i := range fields {
		kvs[fields[i]] = vals[i]
	}

	return kvs, nil
}

func (h *HashmapDataClass) AcquireMapInfo(ctx context.Context, key string) (int64, []string, error) {
	length, err := h.GetRawRedis().HLen(ctx, key).Result()
	if err != nil {
		Log.Errorf(fmt.Sprintf("HashMap HLen failed, err: %s", err))
		return -1, nil, err
	}

	fields, err := h.GetRawRedis().HKeys(ctx, key).Result()

	if err != nil {
		Log.Errorf(fmt.Sprintf("HashMap HLen failed, err: %s", err))
		return -1, nil, err
	}

	return length, fields, nil
}

func (h *HashmapDataClass) IsExist(ctx context.Context, key string, field string) (bool, error) {
	return h.GetRawRedis().HExists(ctx, key, field).Result()
}