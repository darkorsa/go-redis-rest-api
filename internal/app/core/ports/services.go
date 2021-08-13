package ports

import (
	"time"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
)

type QueryService interface {
	Get(key string) (*domain.Item, error)
	GetAll() (*domain.Keys, error)
}

type AuthService interface {
	CreateToken(username string, tokenDuration time.Duration) (*domain.Token, error)
}

type UserService interface {
	Get(username string) (*domain.User, error)
}
