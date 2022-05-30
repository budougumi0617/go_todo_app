package entity

import "time"

type TaskID int64

type Task struct {
	ID      TaskID    `json:"id"`
	Title   string    `json:"title"`
	Status  string    `json:"status" `
	Created time.Time `json:"created"`
}

type Tasks []*Task
