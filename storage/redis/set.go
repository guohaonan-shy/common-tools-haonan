package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type SetDataClass struct {
	cw *CommonRedisWrapper
}

func (s *SetDataClass) GetCommonRedis() *CommonRedisWrapper {
	return s.GetCommonRedis()
}

func (s *SetDataClass) GetRawRedis() *redis.Client {
	return s.cw.RawClient
}

func (s *SetDataClass) NewSetDataClass() *SetDataClass {
	return &SetDataClass{
		cw: NewCommonRedis(CommonRedisClient),
	}
}

func (s *SetDataClass) Add(ctx context.Context, key string, members ...interface{}) (size int64, err error) {
	return s.GetRawRedis().SAdd(ctx, key, members...).Result()
}

// Pop will remove a random elements is set and return this value
func (s *SetDataClass) Pop(ctx context.Context, key string) (val string, err error) {
	return s.GetRawRedis().SPop(ctx, key).Result()
}

// Remove will remove specific elements you assigned in parameters
func (s *SetDataClass) Remove(ctx context.Context, key string, members ...interface{}) (num int64, err error) {
	return s.GetRawRedis().SRem(ctx, key, members...).Result()
}

func (s *SetDataClass) IsMember(ctx context.Context, key string, member interface{}) (yes bool, err error) {
	return s.GetRawRedis().SIsMember(ctx, key, member).Result()
}

// ShowMembers all elements in set of the name user assigned
func (s *SetDataClass) ShowMembers(ctx context.Context, key string) ([]string, error) {
	return s.GetRawRedis().SMembers(ctx, key).Result()
}

// Card return the number of the set
func (s *SetDataClass) Card(ctx context.Context, key string) (num int64, err error) {
	return s.GetRawRedis().SCard(ctx, key).Result()
}

func (s *SetDataClass) Intersection(ctx context.Context, items ...string) ([]string, error) {
	return s.GetRawRedis().SInter(ctx, items...).Result()
}

func (s *SetDataClass) StoreIntersection(ctx context.Context, dest string, items ...string) (int64, error) {
	return s.GetRawRedis().SInterStore(ctx, dest, items...).Result()
}

func (s *SetDataClass) Difference(ctx context.Context, items ...string) ([]string, error)  {
	return s.GetRawRedis().SDiff(ctx, items...).Result()
}

func (s *SetDataClass) StoreDifference(ctx context.Context, dest string, items ...string) (int64, error)  {
	return s.GetRawRedis().SDiffStore(ctx, dest, items...).Result()
}
