package services

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
)

type userService struct {
	repository ports.User
}

func NewUserService(repository ports.User) *userService {
	return &userService{
		repository: repository,
	}
}

func (srv *userService) Get(username string) (*domain.User, error) {
	user, err := srv.repository.Get(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
