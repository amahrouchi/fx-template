package cache

import (
	"fmt"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/redis/go-redis/v9"
)

// CacheContract is the interface for cache services.
type CacheContract interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl int) error
}

// NewCacheService Creates a new CacheContract service.
func NewCacheService(config *fxconfig.Config, logger *fxlogger.Logger) CacheContract {
	// Connect to redis
	host := config.GetString("config.redis.host")
	port := config.GetInt("config.redis.port")
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})

	// Check connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Fatal().Err(err).Msgf("Connection to Redis failed: host: %v, port: %v.", host, port)
	}
	logger.Info().Msgf("Connected successfully to Redis: %v.", pong)

	return &Redis{
		host:   host,
		port:   port,
		client: client,
		logger: logger,
	}
}

// -----------------------------------
// TODO: close the redis connection?
// -----------------------------------
