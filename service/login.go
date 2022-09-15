package service

import (
	"context"
	"fmt"

	"github.com/budougumi0617/go_todo_app/store"
)

type Login struct {
	DB             store.Queryer
	Repo           UserGetter
	TokenGenerator TokenGenerator
}

func (l *Login) Login(ctx context.Context, name, pw string) (string, error) {
	u, err := l.Repo.GetUser(ctx, l.DB, name)
	if err != nil {
		return "", fmt.Errorf("failed to list: %w", err)
	}
	if err := u.ComparePassword(pw); err != nil {
		return "", fmt.Errorf("wrong password: %w", err)
	}
	jwt, err := l.TokenGenerator.GenerateToken(ctx, *u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	return string(jwt), nil
}
