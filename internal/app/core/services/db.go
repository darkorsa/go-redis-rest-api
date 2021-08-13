package services

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
)

type queryService struct {
	repository ports.Repository
}

func NewQueryService(repository ports.Repository) *queryService {
	return &queryService{
		repository: repository,
	}
}

func (srv *queryService) Get(key string) (*domain.Item, error) {
	item, err := srv.repository.Find(key)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (srv *queryService) GetAll() (*domain.Keys, error) {
	items, err := srv.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return items, nil
}
