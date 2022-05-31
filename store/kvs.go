package store

import (
	"context"
	"fmt"
	"time"

	"github.com/budougumi0617/go_todo_app/config"
	"github.com/go-redis/redis/v8"
)

func NewKVS(ctx context.Context, cfg *config.Config) (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
	})
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &KVS{Cli: cli}, nil
}

type KVS struct {
	Cli *redis.Client
}

func (k *KVS) Save(ctx context.Context, key string, userID string) error {
	return k.Cli.Set(ctx, key, userID, 30*time.Minute).Err()
}

func (k *KVS) Load(ctx context.Context, key string) (string, error) {
	return k.Cli.Get(ctx, key).Result()
}
