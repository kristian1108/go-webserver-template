package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-template/src/common"
	"net/http"
)

type Api struct {
	appCache *common.TTLCache[string, interface{}]
}

func New() *Api {
	appCache := common.NewTTL[string, interface{}]()

	return &Api{
		appCache: appCache,
	}
}

func (a *Api) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/test", a.TestHandler)

	return r
}
