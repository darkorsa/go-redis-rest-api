package repository

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

func (r *redisRepository) LRange(key string, start int64, stop int64) (*domain.Item, error) {
	res, err := r.client.LRange(r.ctx, key, start, stop).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Type:  TYPE_LIST,
		Value: res,
	}

	return &item, nil
}

func (r *redisRepository) RPush(key string, values ...interface{}) (int64, error) {
	res, err := r.client.RPush(r.ctx, key, values).Result()

	if err != nil {
		return 0, errors.Wrap(err, "error while rpush to a list")
	}

	return res, err
}

func (r *redisRepository) LPush(key string, values ...interface{}) (int64, error) {
	res, err := r.client.LPush(r.ctx, key, values).Result()

	if err != nil {
		return 0, errors.Wrap(err, "error while lpush to a list")
	}

	return res, err
}

func (r *redisRepository) LRem(key string, count int64, value interface{}) (int64, error) {
	res, err := r.client.LRem(r.ctx, key, count, value).Result()

	if err != nil {
		return 0, errors.Wrap(err, "error while lrem from a list")
	}

	return res, err
}
