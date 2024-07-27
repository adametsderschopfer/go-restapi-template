package app

import (
	"app/internal/lib/http-server/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *App) createRouter() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(app.Log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	//router.Post("/create-url", save.New(app.Log, repository))
	//router.Get("/{alias}", redirect.New(app.Log, repository))

	return router
}
