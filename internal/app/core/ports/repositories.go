package ports

import "github.com/darkorsa/go-redis-http-client/internal/app/core/domain"

type Repository interface {
	Get(key string) (*domain.Item, error)
	List() (*domain.Keys, error)
	Del(key string) (int64, error)
	LRange(key string, start int64, stop int64) (*domain.Item, error)
	RPush(key string, values ...interface{}) (int64, error)
	LPush(key string, values ...interface{}) (int64, error)
	LRem(key string, count int64, value interface{}) (int64, error)
}

type User interface {
	Get(username string) (*domain.User, error)
}
