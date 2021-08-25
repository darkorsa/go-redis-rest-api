package ports

import (
	"time"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
)

type QueryService interface {
	Get(key string) (*domain.Item, error)
	List() (*domain.Keys, error)
	Find(patter string) (*domain.Keys, error)
	Del(key string) (int64, error)
	LRange(key string, start int64, stop int64) (*domain.Item, error)
	RPush(key string, value string) (int64, error)
	LPush(key string, value string) (int64, error)
	LRem(key string, count int64, value string) (int64, error)
}

type AuthService interface {
	CreateToken(username string, tokenDuration time.Duration) (*domain.Token, error)
}

type UserService interface {
	Get(username string) (*domain.User, error)
}
