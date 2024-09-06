package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go-template/src/common"
	"go-template/src/config"
	"log"
	"net/http"
)

type Api struct {
	appCache *common.TTLCache[string, interface{}]
	config   config.TestConfig
}

func New() *Api {
	appCache := common.NewTTL[string, interface{}]()
	appConfig, err := config.Load("src/config")

	if err != nil {
		log.Panic("failed to read config")
	}

	return &Api{
		appCache: appCache,
		config:   appConfig,
	}
}

func (a *Api) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/test", a.TestHandler)

	return r
}
