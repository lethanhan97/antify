package cmd

import (
	"antify/fileserver"
	"antify/health"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func initializeMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.NoCache)
}

func initializeRoutes(router *chi.Mux) {
	router.Get("/health", health.HandleHealth)
	router.Handle("/video/*", fileserver.GetMediaFileServerHandler())
	router.Handle("/*", fileserver.GetHtmlFileServerHandler())
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	initializeMiddlewares(router)
	initializeRoutes(router)

	return router
}
