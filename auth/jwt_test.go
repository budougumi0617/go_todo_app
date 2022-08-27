package auth

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/budougumi0617/go_todo_app/clock"
	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
	"github.com/budougumi0617/go_todo_app/testutil"
	"github.com/budougumi0617/go_todo_app/testutil/fixture"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}
	want = []byte("-----BEGIN PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPrivKey)
	}
}

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
	sut, err := NewJWTer(store, clock.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}
	signed, err := sut.GenerateToken(context.Background(), u)
	if err != nil {
		t.Fatalf("failed to generate jwt: %s", err)
	}
	req := httptest.NewRequest(
		http.MethodGet,
		`https://github.com/budougumi0617`,
		nil)
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

func TestNewJWTer(t *testing.T) {
	got, err := NewJWTer(nil, nil)
	if err != nil {
		t.Fatalf("failed to initialze: %v", err)
	}
	if got.PublicKey == nil {
		t.Error("invalid public key")
	}
	if got.PrivateKey == nil {
		t.Error("invalid private key")
	}
}

func TestJWTer_GenJWT(t *testing.T) {
	ctx := context.Background()
	wantID := entity.UserID(20)
	u := fixture.User(&entity.User{ID: wantID})
	moq := &StoreMock{}
	moq.SaveFunc = func(ctx context.Context, key string, userID entity.UserID) error {
		if userID != wantID {
			t.Errorf("want %d, but got %d", wantID, userID)
		}
		return nil
	}
	sut, err := NewJWTer(moq, clock.RealClocker{})
	if err != nil {
		t.Fatal(err)
	}
	got, err := sut.GenerateToken(ctx, *u)
	if err != nil {
		t.Fatalf("not want err: %v", err)
	}
	if len(got) == 0 {
		t.Errorf("token is empty")
	}
}

func TestJWTer_GetJWT(t *testing.T) {
	t.Parallel()

	c := clock.FixedClocker{}
	want, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/budougumi0617/go_todo_app`).
		Subject("access_token").
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(RoleKey, "test").
		Claim(UserNameKey, "test_user").
		Build()
	if err != nil {
		t.Fatal(err)
	}
	pkey, err := jwk.ParseKey(rawPrivKey, jwk.WithPEM(true))
	if err != nil {
		t.Fatal(err)
	}
	signed, err := jwt.Sign(want, jwt.WithKey(jwa.RS256, pkey))
	if err != nil {
		t.Fatal(err)
	}
	userID := entity.UserID(20)

	ctx := context.Background()
	moq := &StoreMock{}
	moq.LoadFunc = func(ctx context.Context, key string) (entity.UserID, error) {
		return userID, nil
	}
	sut, err := NewJWTer(moq, c)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(
		http.MethodGet,
		`https://github.com/budougumi0617`,
		nil,
	)
	req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, signed))
	got, err := sut.GetToken(ctx, req)
	if err != nil {
		t.Fatalf("want no error, but got %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetToken() got = %v, want %v", got, want)
	}
}

type FixedTomorrowClocker struct{}

func (c FixedTomorrowClocker) Now() time.Time {
	return clock.FixedClocker{}.Now().Add(24 * time.Hour)
}

func TestJWTer_GetJWT_NG(t *testing.T) {
	t.Parallel()

	c := clock.FixedClocker{}
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/budougumi0617/go_todo_app`).
		Subject("access_token").
		IssuedAt(c.Now()).
		Expiration(c.Now().Add(30*time.Minute)).
		Claim(RoleKey, "test").
		Claim(UserNameKey, "test_user").
		Build()
	if err != nil {
		t.Fatal(err)
	}
	pkey, err := jwk.ParseKey(rawPrivKey, jwk.WithPEM(true))
	if err != nil {
		t.Fatal(err)
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, pkey))
	if err != nil {
		t.Fatal(err)
	}

	type moq struct {
		userID entity.UserID
		err    error
	}
	tests := map[string]struct {
		c   clock.Clocker
		moq moq
	}{
		"expire": {
			// トークンのexpire時間より未来の時間を返す。
			c: FixedTomorrowClocker{},
		},
		"notFoundInStore": {
			c: clock.FixedClocker{},
			moq: moq{
				err: store.ErrNotFound,
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			moq := &StoreMock{}
			moq.LoadFunc = func(ctx context.Context, key string) (entity.UserID, error) {
				return tt.moq.userID, tt.moq.err
			}
			sut, err := NewJWTer(moq, tt.c)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(
				http.MethodGet,
				`https://github.com/budougumi0617`,
				nil,
			)
			req.Header.Set(`Authorization`, fmt.Sprintf(`Bearer %s`, signed))
			got, err := sut.GetToken(ctx, req)
			if err == nil {
				t.Errorf("want error, but got nil")
			}
			if got != nil {
				t.Errorf("want nil, but got %v", got)
			}
		})
	}
}
