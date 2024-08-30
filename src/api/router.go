package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-template/src/api/handlers"
	"net/http"
)

type Api struct{}

func New() *Api {
	return &Api{}
}

func (a *Api) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", handlers.TestHandler)

	return r
}
