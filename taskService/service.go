package taskService

import "context"

type Service interface {
	Add(ctx context.Context, task Task) (string, error)
	ChangeStatus(ctx context.Context, id string, status string) (Task, error)
	//GetList(ctx context.Context) ([]Task, error)
	//DeleteTask(ctx context.Context, id string, status string) error
}
