package api

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"io"
	"net/http"
	"sync"
)

type Api struct {
	Router *chi.Mux
}

func New() (*Api, error) {
	api := Api{}
	api.Router = chi.NewRouter()

	api.Router.Use(middleware.RequestID)
	api.Router.Use(middleware.Logger)

	api.Router.Get("/health", api.healthCheck)
	return &api, nil
}

func (s *Api) healthCheck(w http.ResponseWriter, r *http.Request) {
	renderJson(w, r, http.StatusOK, "hello!")
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func renderJson(w http.ResponseWriter, r *http.Request, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		response.Reset()
		bufferPool.Put(response)
	}()
	var err error
	if res != nil {
		err = json.NewEncoder(response).Encode(res)
		if err != nil {
			apiError := ErrorResponse{Message: "Cannot json encode response"}
			_ = json.NewEncoder(response).Encode(apiError)
			statusCode = 500
		}
	}
	w.WriteHeader(statusCode)
	io.Copy(w, response)
}
