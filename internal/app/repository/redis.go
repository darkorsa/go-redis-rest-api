package repository

import (
	"context"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

const (
	TYPE_STRING = "string"
	TYPE_LIST   = "list"
	TYPE_SET    = "set"
	TYPE_ZSET   = "zset"
	TYPE_HASH   = "hash"
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
