package taskService

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db      *sql.DB
	loggger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:      db,
		loggger: log.With(logger, "repo", "sql"),
	}
}

func (rep *repo) Add(ctx context.Context, task Task) (string, error) {
	sql := `
INSERT INTO tasks (id, title, description, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)`

	if task.Title == "" {
		return "", RepoErr
	}

	_, err := rep.db.ExecContext(ctx, sql, task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		return "", RepoErr
	}
	return task.ID, nil
}

func (rep *repo) ChangeStatus(ctx context.Context, id string, status string) (Task, error) {
	var task Task

	sql := `UPDATE tasks
set status = $1, updated_at = NOW()
where id = $2`

	_, err := rep.db.ExecContext(ctx, sql, status, id)
	if err != nil {
		return task, RepoErr
	}

	err = rep.db.QueryRow("select id, title, description, status, created_at, updated_at from tasks where id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	return task, nil
}

//func (rep *repo) GetList(ctx context.Context) ([]Task, error) {
//	var tasks []Task
//
//}

//func(d DB) DeleteTask(ctx context.Context, id string, status string) error{
//
//}
