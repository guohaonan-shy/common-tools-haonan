package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type ListDataClass struct {
	cw *CommonRedisWrapper
}

func (ls *ListDataClass) GetCommonRedis() *CommonRedisWrapper {
	return ls.GetCommonRedis()
}

func (ls *ListDataClass) GetRawRedis() *redis.Client {
	return ls.cw.RawClient
}

func (ls *ListDataClass) NewListDataClass() *ListDataClass {
	return &ListDataClass{
		cw: NewCommonRedis(CommonRedisClient),
	}
}

// Push in fact is an operation of upsert.
// If list exists, we will push in this list.
// Or insert elements after create a new list
func (ls *ListDataClass) Push(ctx context.Context, direction ListOperationDirection, listName string, values ...interface{}) (num int64, err error) {
	switch direction {
	case ListPushFromLeft:
		num, err = ls.GetRawRedis().LPush(ctx, listName, values...).Result()
	case ListPushFromRight:
		num, err = ls.GetRawRedis().RPush(ctx, listName, values...).Result()
	default:
		return 0, errors.New(fmt.Sprintf("Operation type error, operation: %v", direction))
	}
	return
}

func (ls *ListDataClass) Pop(ctx context.Context, direction ListOperationDirection, listName string) (str string, err error) {
	switch direction {
	case ListPushFromLeft:
		str, err = ls.GetRawRedis().LPop(ctx, listName).Result()
	case ListPushFromRight:
		str, err = ls.GetRawRedis().RPop(ctx, listName).Result()
	default:
		return "", errors.New(fmt.Sprintf("Operation type error, operation: %v", direction))
	}
	return
}

// PushX in fact is an operation of insert.
// If list doesn't exist, push will fail.
func (ls *ListDataClass) PushX(ctx context.Context, direction ListOperationDirection, listName string, values ...interface{}) (num int64, err error) {
	switch direction {
	case ListPushFromLeft:
		num, err = ls.GetRawRedis().LPushX(ctx, listName, values...).Result()
	case ListPushFromRight:
		num, err = ls.GetRawRedis().RPushX(ctx, listName, values...).Result()
	default:
		return 0, errors.New(fmt.Sprintf("Operation type error, operation: %v", direction))
	}
	return
}

func (ls *ListDataClass) Range(ctx context.Context, listName string, start, stop int64) ([]string, error) {
	return  ls.GetRawRedis().LRange(ctx, listName, start, stop).Result()
}

func (ls *ListDataClass) Index(ctx context.Context, listName string, index int64) (string, error) {
	return  ls.GetRawRedis().LIndex(ctx, listName, index).Result()
}

func (ls *ListDataClass) Len(ctx context.Context, listName string) (int64, error) {
	return  ls.GetRawRedis().LLen(ctx, listName).Result()
}

// Trim means leave the specific length of list
func (ls *ListDataClass) Trim(ctx context.Context, listName string, start, stop int64) (string, error) {
	return  ls.GetRawRedis().LTrim(ctx, listName, start, stop).Result()
}

func (ls *ListDataClass) ListSetByIndex(ctx context.Context, index int64, key string, value interface{}) (string, error) {
	return ls.GetRawRedis().LSet(ctx, key, index, value).Result()
}

func (ls *ListDataClass) ListInsert(ctx context.Context, operation ListOperationDirection, key string, pivot, value interface{}) (int64, error){
	switch operation {
	case ListInsertFromBefore:
		return ls.GetRawRedis().LInsertBefore(ctx, key, pivot, value).Result()
	case ListInsertFromAfter:
		return ls.GetRawRedis().LInsertAfter(ctx, key, pivot, value).Result()
	default:
		return 0, errors.New(fmt.Sprintf("Operation type error, operation: %v", operation))
	}
}
