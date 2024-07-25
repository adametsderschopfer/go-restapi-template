package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"url-shortener/internal/http-server/middleware/logger"
)

func (app *App) createRouter() http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(logger.New(app.Log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	//router.Post("/create-url", save.New(app.Log, storage))
	//router.Get("/{alias}", redirect.New(app.Log, storage))

	return router
}
