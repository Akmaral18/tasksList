package taskService

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Add          endpoint.Endpoint
	ChangeStatus endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Add:          makeAddEndpoint(s),
		ChangeStatus: makeChangeStatusEndpoint(s),
	}
}

func makeAddEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRequest)
		id, err := s.Add(ctx, req.Task)
		return AddResponse{ID: id, Err: err}, nil
	}
}

func makeChangeStatusEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeStatusRequest)
		taskRes, err := s.ChangeStatus(ctx, req.ID, req.Status)
		return ChangeStatusResponse{Task: taskRes, Err: err}, nil
	}
}
