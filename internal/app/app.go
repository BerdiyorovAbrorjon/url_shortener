package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BerdiyorovAbrorjon/url-shortener/config"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/rediscache"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/service"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/transport/http"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/database/postgres"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/database/redis"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/httpserver"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/logger"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/token"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("app - Run - config.NewConfig: %w", err)
	}

	l := logger.New(cfg.LogLevel)

	// Initialize Postgres Pool
	dbPool, err := postgres.NewPool(cfg.DbSource)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.NewPool: %w", err))
	}
	l.Info("Connected to postgres...")

	// Initialize Redis
	rdb, err := redis.NewClient(cfg.RedisAddress, "", 0)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis.NewClient: %w", err))
	}
	l.Info("Connected to redis...")

	//DB & Cache
	pgStore := pgstore.New(dbPool)
	redisCache := rediscache.NewRedisCache(rdb, time.Hour)

	//Repo
	userRepo := repository.NewUserRepository(pgStore)
	urlRepo := repository.NewUrlRepository(pgStore, redisCache)

	//Services
	userService := service.NewUserService(userRepo)
	urlService := service.NewUrlService(urlRepo)

	service := service.Service{
		User: userService,
		Url:  urlService,
	}

	//Token make JWT
	tokenMaker, err := token.NewJWTMaker(cfg.TokenSymmetricKey)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - createTokenMaker: %w", err))
	}

	//HTTP server
	handler := http.NewHandler(cfg, &service, tokenMaker)

	router := handler.InitRouter()

	httpServer := httpserver.New(router, httpserver.Port(cfg.HttpServerPort))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
