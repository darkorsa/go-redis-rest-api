package services

import (
	"time"

	"github.com/darkorsa/go-redis-http-client/internal/app/core/domain"
	"github.com/darkorsa/go-redis-http-client/internal/pkg/token"
)

type authService struct {
	tokenMaker token.Maker
}

func NewAuthService(tokenMaker token.Maker) *authService {
	return &authService{
		tokenMaker: tokenMaker,
	}
}

func (srv *authService) CreateToken(username string, tokenDuration time.Duration) (*domain.Token, error) {
	accessToken, err := srv.tokenMaker.CreateToken(
		username,
		tokenDuration,
	)

	if err != nil {
		return nil, err
	}

	token := &domain.Token{
		AccessToken: accessToken,
		ExpiresAt:   time.Now().Add(tokenDuration),
	}

	return token, nil
}
