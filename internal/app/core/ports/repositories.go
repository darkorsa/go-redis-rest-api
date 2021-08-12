package ports

import "github.com/darkorsa/go-redis-http-client/internal/app/core/domain"

type Repository interface {
	Find(key string) (*domain.Item, error)
	FindAll() ([]*domain.Key, error)
}

type User interface {
	Get(username string) (*domain.User, error)
}
