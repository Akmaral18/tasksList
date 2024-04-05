package taskService

import (
	"context"
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"-" sql:"created_at"`
	UpdatedAt   time.Time `json:"-"`
}
type Repository interface {
	Add(ctx context.Context, task Task) (string, error)
	ChangeStatus(ctx context.Context, id string, status string) (Task, error)
	//GetList(ctx context.Context) ([]Task, error)
	//DeleteTask(ctx context.Context, id string, status string) error
}
