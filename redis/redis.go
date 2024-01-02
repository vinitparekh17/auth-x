package redis

import (
	"context"
	"errors"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"github.com/vinitparekh17/project-x/handler"
)

var ctx = context.Background()
var rdb *redis.Client

func Init() {
	// Initialize redis client with default options
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "secret",
		DB:       0,
	})

	// Ping redis server to check connection
	status := rdb.Ping(ctx)
	slog.Info("Redis pinging... " + status.String())
}

func Set(key string, value string) {
	err := rdb.Set(ctx, key, value, 0).Err()
	handler.ErrorHandler(err)
}

func Get(key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		handler.ErrorHandler(errors.New("key does not exist"))
	} else if err != nil {
		handler.ErrorHandler(err)
	} else {
		return val
	}
	return ""
}
