package repository

import (
	"fmt"

	"github.com/darkorsa/go-redis-rest-api/internal/app/core/domain"
	"github.com/darkorsa/go-redis-rest-api/internal/app/core/ports"
	"github.com/darkorsa/go-redis-rest-api/internal/app/util"
)

type userRepository struct {
	config *util.Config
}

func NewUserRepository(config *util.Config) (ports.User, error) {
	repo := &userRepository{
		config: config,
	}

	return repo, nil
}

func (r *userRepository) Get(username string) (*domain.User, error) {
	user := domain.NewUser(r.config.AuthUsername, r.config.AuthPassword)

	if username != user.Username {
		return nil, fmt.Errorf("no user with username: %s", username)
	}

	return user, nil
}
