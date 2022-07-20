package service

import (
	"context"
	"testing"

	"github.com/budougumi0617/go_todo_app/auth"

	"github.com/google/go-cmp/cmp"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

func TestAddTask_AddTask(t *testing.T) {
	t.Parallel()

	wantUID := entity.UserID(10)
	wantTitle := "test title"
	wantTask := &entity.Task{
		UserID: wantUID,
		Title:  wantTitle,
		Status: entity.TaskStatusTodo,
	}
	type TaskAdderMockParameter struct {
		in  *entity.Task
		err error
	}
	tests := map[string]struct {
		title string
		want  *entity.Task
		taprm TaskAdderMockParameter
	}{
		"正常系": {
			title: wantTitle,
			want:  wantTask,
			taprm: TaskAdderMockParameter{
				in:  wantTask,
				err: nil,
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ctx := auth.SetUserID(context.Background(), 10)

			moqDB := &ExecerMock{}
			moqRepo := &TaskAdderMock{}
			moqRepo.AddTaskFunc = func(pctx context.Context, db store.Execer, task *entity.Task) error {
				if ctx != pctx {
					t.Fatalf("not want context %v", pctx)
				}
				if db != moqDB {
					t.Fatalf("not want db %v", db)
				}
				if d := cmp.Diff(task, tt.taprm.in); len(d) != 0 {
					t.Fatalf("differs: (-got +want)\n%s", d)
				}
				return tt.taprm.err
			}
			a := &AddTask{
				DB:   moqDB,
				Repo: moqRepo,
			}
			got, err := a.AddTask(ctx, tt.title)
			if err != nil {
				t.Fatalf("want no error, but got %v", err)
				return
			}
			if d := cmp.Diff(got, tt.want); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
