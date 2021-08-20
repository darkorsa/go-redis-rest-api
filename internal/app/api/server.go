package api

import (
	"github.com/darkorsa/go-redis-http-client/docs"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/ports"
	"github.com/darkorsa/go-redis-http-client/internal/app/core/services"
	"github.com/darkorsa/go-redis-http-client/internal/app/repository"
	"github.com/darkorsa/go-redis-http-client/internal/app/util"
	apiErrors "github.com/darkorsa/go-redis-http-client/internal/pkg/api-errors"
	"github.com/darkorsa/go-redis-http-client/internal/pkg/token"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func (s *Server) StartServer() {
	r := gin.Default()

	r.POST("/token", s.GetToken)

	authorized := r.Group("/")
	authorized.Use(authMiddleware(s.tokenMaker))
	{
		authorized.GET("/keys/:id", s.GetKey)
		authorized.GET("/keys", s.GetKeys)
		authorized.DELETE("/keys/:id", s.DelKey)
		authorized.GET("/list/key/:id", s.ListGet)
		authorized.DELETE("/list/key/:id", s.ListDel)
		authorized.POST("/list/rpush/key/:id", s.ListRPush)
		authorized.POST("/list/lpush/key/:id", s.ListLPush)
	}

	docs.SwaggerInfo.Title = "Redis API"
	docs.SwaggerInfo.Description = "API providing access to Redis database."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = s.config.ServerAddress
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(s.config.ServerAddress); err != nil {
		panic(err)
	}
}

func (s *Server) badRequestError(msg string, c *gin.Context) {
	apiErr := apiErrors.NewBadRequestError(msg)
	c.JSON(apiErr.GetStatus(), apiErr)
}

func (s *Server) internalServerError(mgs string, c *gin.Context) {
	apiErr := apiErrors.NewInternalServerError(mgs)
	c.JSON(apiErr.GetStatus(), apiErr)
}

func (s *Server) notFoundError(msg string, c *gin.Context) {
	apiErr := apiErrors.NewNotFoundError(msg)
	c.JSON(apiErr.GetStatus(), apiErr)
}
