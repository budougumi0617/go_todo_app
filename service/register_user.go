package service

import (
	"context"
	"fmt"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	DB   store.Execer
	Repo UserRegister
}

func (r *RegisterUser) RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot encrypt password: %w", err)
	}
	t := &entity.User{
		Name:     name,
		Password: string(pw),
		Role:     role,
	}

	if err := r.Repo.RegisterUser(ctx, r.DB, t); err != nil {
		return nil, fmt.Errorf("failed to register: %w", err)
	}
	return t, nil
}
