package app

import (
	"app/internal/lib/http-server/middleware/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
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

	/*Example routes*/
	//router.Post("/create-url", save.New(app.Log, app.repository))
	//router.Get("/{alias}", redirect.New(app.Log, app.repository))
	router.Get("/", testMethod(app.Log))
	return router
}

func testMethod(log *slog.Logger) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Got URL")

		render.HTML(writer, request, "<h1>Hello world</h1>")
	}
}
