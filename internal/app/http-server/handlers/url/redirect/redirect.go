package redirect

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	resp "url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage"
)

type UrlGetter interface {
	GetUrlByAlias(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "handlers.url.Redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())),
		)

		alias := chi.URLParam(request, "alias")
		if len(alias) == 0 {
			log.Info("alias is empty")
			render.JSON(writer, request, resp.Error("Invalid request"))
			return
		}

		resUrl, err := urlGetter.GetUrlByAlias(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Info("Url not found", sl.Err(err))
				render.JSON(writer, request, resp.Error("Url not found"))
				return
			}

			log.Error("Failed to get url", sl.Err(err))
			render.JSON(writer, request, resp.Error("Internal error"))
			return
		}

		log.Info("Got URL", slog.String("URL", resUrl))

		http.Redirect(writer, request, resUrl, http.StatusFound)
	}
}
