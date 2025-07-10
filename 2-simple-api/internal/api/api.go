package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/lapeko/go__test-practice/2-simple-api/internal/api/domain/health"
	"github.com/lapeko/go__test-practice/2-simple-api/internal/api/middleware"
	"github.com/lapeko/go__test-practice/2-simple-api/internal/logger"
	"net/http"
	"os"
)

type api struct {
	r *chi.Mux
}

type Api interface {
	SetupMiddlewares()
	SetupRouter()
	Start() error
}

func New() Api {
	logger.Log.Debug("Api. Instance creation")
	return &api{chi.NewRouter()}
}

func (a *api) SetupMiddlewares() {
	a.r.Use(middleware.LogMiddleware)
	logger.Log.Debug("Api. Setting up middlewares")
}

func (a *api) SetupRouter() {
	logger.Log.Debug("Api. Setting up router")
	a.r.Mount("/health", health.GetController())
	//a.r.Mount("/orders", order.GetController())
	//a.r.Mount("/users", user.GetController())
}

func (a *api) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logger.Log.Debugf("Api. Running on: %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), a.r)
}
