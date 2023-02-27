package cache

import (
	"context"
	"tiny-url/config"

	"github.com/go-redis/redis/v8"
)
type CacheInterface interface {
	InsertCache(ctx context.Context,shorUrl,longUrl string) error
	GetCache(ctx context.Context, shorUrl string) (string, error)
}

type cache struct {
	rdb *redis.Client
}
func NewCache (rdb *redis.Client) CacheInterface {
	return &cache{
		rdb: rdb,
	}
}

func (c *cache) InsertCache(ctx context.Context,shorUrl,longUrl string) error {
	err:= c.rdb.Set(ctx, shorUrl, longUrl, config.Cfg.Redis.TTL).Err()
	if err != nil {
		return err
	}
	return nil
}
func (c *cache) GetCache(ctx context.Context, shorUrl string) (string, error) {
	longUrl, err := c.rdb.Get(ctx, shorUrl).Result()
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
