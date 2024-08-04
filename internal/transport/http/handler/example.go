package handler

import (
	"app/internal/service"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func (h *Handlers) exampleHandler(log *slog.Logger, exampleService service.Example) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Info("Got URL")

		exampleService.Create("example name")

		render.HTML(writer, request, "<h1>Hello world</h1>")
	}
}
