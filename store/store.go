package store

import (
	"errors"

	"github.com/budougumi0617/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{tasks: map[int]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	lastID int
	tasks  map[int]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (int, error) {
	ts.lastID++
	t.ID = ts.lastID
	ts.tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) Get(id int) (*entity.Task, error) {
	if ts, ok := ts.tasks[id]; ok {
		return ts, nil
	}
	return nil, ErrNotFound
}
