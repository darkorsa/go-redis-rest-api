package api

import (
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/services"
	"github.com/darkorsa/go-redis-http-client/internal/app/repository"
	"github.com/darkorsa/go-redis-http-client/internal/app/util"
	"github.com/darkorsa/go-redis-http-client/internal/pkg/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config       *util.Config
	tokenMaker   token.Maker
	queryService ports.QueryService
	authService  ports.AuthService
	userService  ports.UserService
}

func NewServer(config *util.Config) (*Server, error) {
	redisRepo, err := repository.NewRedisRepository(
		config.RedisServer,
		config.RedisPort,
		config.RedisPaswword,
		config.RedisDB,
	)
	if err != nil {
		panic(err)
	}
	userRepo, err := repository.NewUserRepository(config)
	if err != nil {
		panic(err)
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		panic(err)
	}

	queryService := services.NewQueryService(redisRepo)
	authService := services.NewAuthService(tokenMaker)
	userService := services.NewUserService(userRepo)

	server := &Server{
		config:       config,
		tokenMaker:   tokenMaker,
		queryService: queryService,
		authService:  authService,
		userService:  userService,
	}

	return server, nil
}

func (s *Server) StartServer() {
	r := gin.Default()

	r.POST("/token", s.GetToken)

	authorized := r.Group("/")
	authorized.Use(authMiddleware(s.tokenMaker))
	{
		authorized.GET("/keys/:id", s.GetKey)
		authorized.GET("/keys", s.GetKeys)
		authorized.DELETE("/keys/:id", s.DelKey)
	}

	if err := r.Run(s.config.ServerAddress); err != nil {
		panic(err)
	}
}
