package taskService

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type (
	AddRequest struct {
		Task Task `json:"task"`
	}

	AddResponse struct {
		ID  string `json:"id"`
		Err error  `json:"error,omitempty"`
	}

	ChangeStatusRequest struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}

	ChangeStatusResponse struct {
		Task Task  `json:"task"`
		Err  error `json:"error,omitempty"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeAddReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req AddRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err1 := json.Unmarshal(body, &req)
	if err1 != nil {
		return nil, err1
	}
	return req, nil
}

func decodeChangeStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ChangeStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
