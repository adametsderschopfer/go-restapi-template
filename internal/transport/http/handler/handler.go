package handler

import (
	"app/internal/service"
	"app/pkg/http-server/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"net/http"
)

type Handlers struct {
	log      *slog.Logger
	services *service.Services
}

func NewTransportHandlers(log *slog.Logger, services *service.Services) http.Handler {
	handlers := Handlers{
		log:      log,
		services: services,
	}

	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(handlers.log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	/*Routes*/
	router.Get("/", handlers.exampleHandler(handlers.log, handlers.services.Example))

	return router
}
