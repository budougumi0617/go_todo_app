package service

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}
type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer, id entity.UserID) (entity.Tasks, error)
}
type UserRegister interface {
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}
