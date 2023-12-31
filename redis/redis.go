package redis

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func Init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "secret",
		DB:       0,
	})

	status := rdb.Ping(ctx)
	slog.Info("Redis pinging... " + status.String())
}
