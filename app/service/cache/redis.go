package cache

import (
	"context"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Redis cache service
type Redis struct {
	host   string
	port   int
	client *redis.Client
	logger *fxlogger.Logger
}

// Get a key from the redis cache
func (r *Redis) Get(key string) (string, error) {
	return "todo", nil
}

// Set a key into the redis cache
func (r *Redis) Set(key string, value string) error {
	return nil
}
