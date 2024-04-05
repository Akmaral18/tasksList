package taskService

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"time"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, logger log.Logger) *service {
	return &service{repository: rep, logger: logger}
}

func (s service) Add(ctx context.Context, task Task) (string, error) {
	logger := log.With(s.logger, "method", "Add")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	task.ID = id
	task.Status = "ToDo"
	task.CreatedAt = time.Now().UTC()

	if _, err := s.repository.Add(ctx, task); err != nil {
		level.Error(logger).Log("err", err)
		return "", errors.Errorf("error adding task - %w", err)
	}
	return id, nil

}

func (s service) ChangeStatus(ctx context.Context, id string, status string) (Task, error) {
	logger := log.With(s.logger, "method", "ChangeStatus")
	task, err := s.repository.ChangeStatus(ctx, id, status)
	if err != nil {
		level.Error(logger).Log("err", err)
		return task, errors.Errorf("error changing status - %w", err)
	}
	return task, nil
}
