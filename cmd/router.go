package cmd

import (
	"musicserver/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func initializeMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
}

func initializeRoutes(router *chi.Mux) {
	router.Get("/health", health.HandleHealth)

}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	initializeMiddlewares(router)
	initializeRoutes(router)

	return router
}
