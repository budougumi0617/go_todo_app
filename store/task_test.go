package store

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/budougumi0617/go_todo_app/clock"
	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/testutil"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// entity.Taskを作成する他のテストケースと混ざるとテストがフェイルする。
	// そのため、トランザクションをはることでこのテストケースの中だけのテーブル状態にする。
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテスト毛＝すが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	// 一度きれいにしておく
	if _, err := tx.ExecContext(ctx, "DELETE FROM task;"); err != nil {
		t.Logf("failed to initialize task: %v", err)
	}

	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 2", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 3", Status: "done",
			Created: c.Now(), Modified: c.Now(),
		},
	}
	result, err := tx.ExecContext(ctx,
		`INSERT INTO task (title, status, created, modified)
    		VALUES
    		    (?, ?, ?, ?),
    		    (?, ?, ?, ?),
    		    (?, ?, ?, ?);`,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)
	t.Cleanup(func() {
		if _, err := tx.ExecContext(ctx, "DELETE FROM task WHERE id IN (?, ?, ?);", wants[0].ID, wants[1].ID, wants[2].ID); err != nil {
			t.Logf("failed to delete task: %v", err)
		}
	})

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexected error: %v", err)
	}
	if d := cmp.Diff(gots, wants, cmpopts.EquateApproxTime(1*time.Second)); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func TestRepository_AddTask(t *testing.T) {
	ctx := context.Background()

	c := clock.FixedClocker{}
	wantID := 20
	okTask := &entity.Task{
		Title:    "ok task",
		Status:   "todo",
		Created:  c.Now(),
		Modified: c.Now(),
	}

	tests := map[string]struct {
		task    *entity.Task
		wantErr bool
	}{
		"ok": {
			task: okTask,
		},
	}
	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() { db.Close() })
			mock.ExpectExec(
				// エスケープが必要
				`INSERT INTO task \(title, status, created, modified\) VALUES \(\?, \?, \?, \?, \?\)`,
			).WithArgs(tt.task.Title, tt.task.Status, c.Now(), c.Now()).
				WillReturnResult(sqlmock.NewResult(int64(wantID), 1))

			xdb := sqlx.NewDb(db, "mysql")
			r := &Repository{Clocker: c}
			if err := r.AddTask(ctx, xdb, tt.task); (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
