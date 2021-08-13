package repository

import (
	"context"
	"fmt"
	"reflect"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func newRedisClient(server string, port string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     server + ":" + port,
		Password: password,
		DB:       db,
	})

	return client, nil
}

func NewRedisRepository(server string, port string, password string, db int) (ports.Repository, error) {
	redisClient, err := newRedisClient(server, port, password, db)
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	repo := &redisRepository{
		client: redisClient,
		ctx:    context.Background(),
	}

	return repo, nil
}

func (r *redisRepository) Find(key string) (*domain.Item, error) {
	t, err := r.client.Type(r.ctx, key).Result()

	if err != nil {
		return nil, errors.Wrap(err, "unable to get key value type")
	}

	if t == "string" {
		return r.fetchString(key)
	} else if t == "list" {
		return r.fetchList(key)
	} else if t == "set" {
		return r.fetchSet(key)
	} else if t == "none" {
		return nil, nil
	} else {
		return nil, fmt.Errorf("type: %s is not supported", t)
	}
}

func (r *redisRepository) FindAll() (*domain.Keys, error) {
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

	var val []string

	item := domain.Item{
		Key:   key,
		Value: append(val, res),
	}

	return &item, nil
}

func (r *redisRepository) fetchList(key string) (*domain.Item, error) {
	res, err := r.client.LRange(r.ctx, key, 0, -1).Result()

	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, err
	}

	item := domain.Item{
		Key:   key,
		Value: res,
	}

	return &item, nil
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
		Value: res,
	}

	return &item, nil
}
