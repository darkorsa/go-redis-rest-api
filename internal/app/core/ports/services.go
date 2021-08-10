package ports

import "github.com/darkorsa/go-redis-http-client/internal/app/core/domain"

type Service interface {
	Get(key string) (*domain.Item, error)
	GetAll() ([]*domain.Key, error)
}
