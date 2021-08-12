package repository

import (
	"context"
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
	val, err := r.client.Get(r.ctx, key).Result()
	switch {
	case err == redis.Nil:
		return nil, nil
	case err != nil:
		return nil, errors.Wrap(err, "GET command failed")
	}

	item := domain.Item{
		Key:   key,
		Value: val,
	}

	return &item, nil
}

func (r *redisRepository) FindAll() ([]*domain.Key, error) {
	res, err := r.client.Do(r.ctx, "KEYS", "*").Result()

	if err != nil {
		return nil, errors.Wrap(err, "KEYS * command failed")
	}

	s := reflect.ValueOf(res)

	items := []*domain.Key{}
	for i := 0; i < s.Len(); i++ {
		item := domain.Key{
			Value: s.Index(i).Interface().(string),
		}
		items = append(items, &item)
	}

	return items, nil
}
