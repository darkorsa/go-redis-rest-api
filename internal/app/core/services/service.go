package services

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
)

type service struct {
	repository ports.Repository
}

func NewService(repository ports.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (srv *service) Get(key string) (*domain.Item, error) {
	item, err := srv.repository.Find(key)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (srv *service) GetAll() ([]*domain.Key, error) {
	items, err := srv.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}
