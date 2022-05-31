package auth

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
	"github.com/budougumi0617/go_todo_app/testutil"
)

// デバッグコード
func TestJWTer(t *testing.T) {
	store := &store.KVS{Cli: testutil.OpenRedisForTest(t)}
	wantID := entity.UserID(20)
	u := entity.User{
		ID:       wantID,
		Name:     "budougumi",
		Password: "test",
		Role:     "admin",
		Created:  time.Time{},
		Modified: time.Time{},
	}
	sut, err := NewJWTer(store)
	if err != nil {
		t.Fatal(err)
	}
	signed, err := sut.GenJWT(context.Background(), u)
	if err != nil {
		t.Fatalf("failed to generate jwt: %s", err)
	}
	req, _ := http.NewRequest(http.MethodGet, `https://github.com/budougumi0617`, nil)
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, signed))
	t.Logf("generated\n%s\n", signed)

	// HTTPハンドラーが受け取ったリクエストを想定
	req, err = sut.FillContext(req)
	if err != nil {
		t.Fatalf("failed to initialize request: %v", err)
	}
	if !IsAdmin(req.Context()) {
		t.Error("IsAdmin() should be true")
	}
	got, ok := GetUserID(req.Context())
	if !ok {
		t.Fatal("GetUserID() should be true")
	}
	if got != wantID {
		t.Errorf("want %d, but got %d", wantID, got)
	}
}
