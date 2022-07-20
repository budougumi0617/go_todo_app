package entity

import "time"

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Status  string    `json:"status" `
	Created time.Time `json:"created"`
}

type Tasks []*Task
