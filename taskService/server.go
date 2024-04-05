package taskService

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/task").Handler(httptransport.NewServer(
		endpoints.Add, decodeAddReq, encodeResponse))

	r.Methods("POST").Path("/task/changeStatus").Handler(httptransport.NewServer(
		endpoints.ChangeStatus, decodeChangeStatusReq, encodeResponse))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
