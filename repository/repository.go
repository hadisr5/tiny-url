package repository

import (
	"database/sql"
	"tiny-url/repository/cache"
	"tiny-url/repository/database"

	"github.com/go-redis/redis/v8"
)

type Repository struct {
	Database database.DatabaseInterface
	Cache    cache.CacheInterface
}

func NewRepository(mdb *sql.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Database: database.NewDatabase(mdb),
		Cache:    cache.NewCache(rdb),
	}
}
