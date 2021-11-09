package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type SortedSetDataClass struct {
	cw *CommonRedisWrapper
}

func (s *SortedSetDataClass) SortedSetDataClass() *CommonRedisWrapper {
	return s.cw
}

func (s *SortedSetDataClass) GetRawRedis() *redis.Client {
	return s.cw.RawClient
}

func (s *SortedSetDataClass) NewSortedSetDataClass() *SortedSetDataClass {
	return &SortedSetDataClass{
		cw: NewCommonRedis(CommonRedisClient),
	}
}

func (s *SortedSetDataClass) NewZNodeList() []*redis.Z  {
	list := make([]*redis.Z, 0)
	return list
}

func (s *SortedSetDataClass) GenerateZNode(element interface{}, score float64) *redis.Z {
	return &redis.Z{
		Member: element,
		Score: score,
	}
}

func (s *SortedSetDataClass) AppendNode(dst []*redis.Z, src ...*redis.Z) []*redis.Z  {
	return append(dst, src...)
}

func (s *SortedSetDataClass) GenerateRangeBy(minScore, maxScore string, offset, count int64) *redis.ZRangeBy  {
	return &redis.ZRangeBy{
		Min: minScore,
		Max: maxScore,
		Offset: offset,
		Count: count,
	}
}

// Add updates elements' score if elements exist
func (s *SortedSetDataClass) Add(ctx context.Context, key string, elements []*redis.Z) (int64, error)  {
	return s.GetRawRedis().ZAdd(ctx, key, elements...).Result()
}

func (s *SortedSetDataClass) Range(ctx context.Context, key string, from, to int64) ([]string, error) {
	return s.GetRawRedis().ZRange(ctx, key, from, to).Result()
}

func (s *SortedSetDataClass) IncrBy(ctx context.Context, key string, score float64, element string) (float64, error) {
	return s.GetRawRedis().ZIncrBy(ctx, key, score, element).Result()
}

func (s *SortedSetDataClass) Card(ctx context.Context, key string) (int64, error) {
	return s.GetRawRedis().ZCard(ctx, key).Result()
}

func (s *SortedSetDataClass) Count(ctx context.Context, key, left, right string) (int64, error)  {
	return s.GetRawRedis().ZCount(ctx, key, left, right).Result()
}

func (s *SortedSetDataClass) RangeByScore(ctx context.Context, key, minScore, maxScore string, offset, count int64) ([]string, error) {
	return s.GetRawRedis().ZRangeByScore(ctx, key, s.GenerateRangeBy(minScore, maxScore, offset, count)).Result()
}

func (s *SortedSetDataClass) Rank(ctx context.Context, key, element string) (int64, error)  {
	return s.GetRawRedis().ZRank(ctx, key, element).Result()
}

func (s *SortedSetDataClass) Score(ctx context.Context, key, element string) (float64, error)  {
	return s.GetRawRedis().ZScore(ctx, key, element).Result()
}

func (s *SortedSetDataClass) Delete(ctx context.Context, key string, elements ...interface{}) (int64, error) {
	return s.GetRawRedis().ZRem(ctx, key, elements).Result()
}

// DeleteByRank sort by increment.
// Index, which is negative, will remove elements from the highest to lowest
func (s *SortedSetDataClass) DeleteByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	return s.GetRawRedis().ZRemRangeByRank(ctx, key, start, stop).Result()
}