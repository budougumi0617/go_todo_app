package main

import (
	"net/http"

	"github.com/budougumi0617/go_todo_app/handler"
	"github.com/budougumi0617/go_todo_app/store"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	mux.Handle("/tasks", &handler.AddTask{Store: store.Tasks, Validator: v})
	return mux
}
