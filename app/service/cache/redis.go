package cacheService

import (
	"context"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/redis/go-redis/v9"
	"time"
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
	result, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		r.logger.Err(err).Msgf("Unable to retrieve key=%s", key)
		return "", err
	}

	return result, nil
}

// Set a key into the redis cache
func (r *Redis) Set(key string, value string, ttl int) error {
	err := r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		r.logger.Err(err).
			Str("key", key).
			Str("value", value).
			Int("ttl", ttl).
			Msgf("Unable to set data into Redis.")
		return err
	}

	return nil
}
