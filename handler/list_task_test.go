package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
	"github.com/budougumi0617/go_todo_app/testutil"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		tasks map[int]*entity.Task
		want  want
	}{
		"ok": {
			tasks: map[int]*entity.Task{
				1: {
					ID:     1,
					Title:  "test1",
					Status: "todo",
				},
				2: {
					ID:     2,
					Title:  "test2",
					Status: "done",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json.golden",
			},
		},
		"empty": {
			tasks: map[int]*entity.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			sut := ListTask{Store: &store.TaskStore{Tasks: tt.tasks}}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
