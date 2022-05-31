package store

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/budougumi0617/go_todo_app/entity"

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

func (k *KVS) Save(ctx context.Context, key string, userID entity.UserID) error {
	id := int64(userID)
	return k.Cli.Set(ctx, key, id, 30*time.Minute).Err()
}

func (k *KVS) Load(ctx context.Context, key string) (entity.UserID, error) {
	result, err := k.Cli.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("%q is not to number :%w", result, err)
	}
	return entity.UserID(id), nil
}
