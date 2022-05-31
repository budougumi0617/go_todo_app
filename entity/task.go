package entity

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID       TaskID     `json:"id" db:"id"`
	UserID   UserID     `json:"user_id" db:"user_id"`
	Title    string     `json:"title" db:"title"`
	Status   TaskStatus `json:"status" db:"status"`
	Created  time.Time  `json:"created" db:"created"`
	Modified time.Time  `json:"modified" db:"modified"`
}

type Tasks []*Task
