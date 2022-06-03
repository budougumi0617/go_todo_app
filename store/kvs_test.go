package store

import (
	"context"
	"testing"
	"time"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/testutil"
)

func TestKVS_Save(t *testing.T) {
	cli := testutil.OpenRedisForTest(t)

	sut := &KVS{Cli: cli}
	key := "TestKVS_Save"
	uid := entity.UserID(1234)
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestKVS_Load(t *testing.T) {
	cli := testutil.OpenRedisForTest(t)

	t.Run("ok", func(t *testing.T) {
		sut := &KVS{Cli: cli}
		key := "TestKVS_Load_ok"
		uid := entity.UserID(1234)
		ctx := context.Background()
		cli.Set(ctx, key, int64(uid), 30*time.Minute)
		t.Cleanup(func() {
			cli.Del(ctx, key)
		})
		got, err := sut.Load(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if got != uid {
			t.Errorf("want %d, but got %d", uid, got)
		}
	})

	t.Run("notFound", func(t *testing.T) {
		sut := &KVS{Cli: cli}
		key := "TestKVS_Save_notFound"
		ctx := context.Background()
		if got, err := sut.Load(ctx, key); err == nil {
			t.Errorf("want no error, but got %v(value = %d)", err, got)
		}
	})
}