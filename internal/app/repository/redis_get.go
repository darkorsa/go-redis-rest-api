package repository

import (
	"fmt"
	"reflect"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func (r *redisRepository) Get(key string) (*domain.Item, error) {
	t, err := r.client.Type(r.ctx, key).Result()

	if err != nil {
		return nil, errors.Wrap(err, "unable to get key value type")
	}

	switch {
	case t == TYPE_STRING:
		return r.fetchString(key)
	case t == TYPE_LIST:
		return r.fetchList(key)
	case t == TYPE_SET:
		return r.fetchSet(key)
	case t == TYPE_ZSET:
		return r.fetchSortedSet(key)
	case t == TYPE_HASH:
		return r.fetchHash(key)
	case t == "none":
		return nil, nil
	default:
		return nil, fmt.Errorf("type: %s is not supported", t)
	}
}

func (r *redisRepository) List() (*domain.Keys, error) {
	res, err := r.client.Do(r.ctx, "KEYS", "*").Result()

	if err != nil {
		return nil, errors.Wrap(err, "KEYS * command failed")
	}

	s := reflect.ValueOf(res)

	var items []string
	for i := 0; i < s.Len(); i++ {
		items = append(items, s.Index(i).Interface().(string))
	}

	keys := domain.Keys{
		Keys: items,
	}

	return &keys, nil
}

func (r *redisRepository) fetchString(key string) (*domain.Item, error) {
	res, err := r.client.Get(r.ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Type:  TYPE_STRING,
		Value: res,
	}

	return &item, nil
}

func (r *redisRepository) fetchList(key string) (*domain.Item, error) {
	return r.LRange(key, 0, -1)
}

func (r *redisRepository) fetchSet(key string) (*domain.Item, error) {
	res, err := r.client.SMembers(r.ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Type:  TYPE_SET,
		Value: res,
	}

	return &item, nil
}

func (r *redisRepository) fetchSortedSet(key string) (*domain.Item, error) {
	res, err := r.client.ZRangeByScore(r.ctx, key, &redis.ZRangeBy{}).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Type:  TYPE_ZSET,
		Value: res,
	}

	return &item, nil
}

func (r *redisRepository) fetchHash(key string) (*domain.Item, error) {
	res, err := r.client.HGetAll(r.ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Type:  TYPE_HASH,
		Value: res,
	}

	return &item, nil
}
