package ports

import "github.com/darkorsa/go-redis-http-client/internal/app/core/domain"

type Repository interface {
	Get(key string) (*domain.Item, error)
	GetAll() (*domain.Keys, error)
	Del(key string) (int64, error)
}

type User interface {
	Get(username string) (*domain.User, error)
}
