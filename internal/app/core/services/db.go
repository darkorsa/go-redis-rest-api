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
	return srv.repository.Get(key)
}

func (srv *queryService) List() (*domain.Keys, error) {
	return srv.repository.List()
}

func (srv *queryService) Find(pattern string) (*domain.Keys, error) {
	return srv.repository.Find(pattern)
}

func (srv *queryService) Del(key string) (int64, error) {
	return srv.repository.Del(key)
}

func (srv *queryService) LRange(key string, start int64, stop int64) (*domain.Item, error) {
	return srv.repository.LRange(key, start, stop)
}

func (srv *queryService) RPush(key string, value string) (int64, error) {
	return srv.repository.RPush(key, []byte(value))
}

func (srv *queryService) LPush(key string, value string) (int64, error) {
	return srv.repository.LPush(key, []byte(value))
}

func (srv *queryService) LRem(key string, count int64, value string) (int64, error) {
	return srv.repository.LRem(key, count, value)
}
