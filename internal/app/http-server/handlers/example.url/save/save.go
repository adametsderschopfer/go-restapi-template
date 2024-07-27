package save

import (
	resp "app/internal/lib/api/response"
	"app/internal/lib/logger/sl"
	"app/internal/lib/random"
	"app/internal/repository"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type UrlSaver interface {
	SaveUrl(urlToSave string, alias string) (int64, error)
}

// todo: move to config
const aliasLength = 8

func New(log *slog.Logger, urlSaver UrlSaver) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "handlers.url.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(request.Context())),
		)

		var reqBody Request
		err := render.DecodeJSON(request.Body, &reqBody)
		if err != nil {
			log.Error("Failed to parse request body", sl.Err(err))
			render.JSON(writer, request, resp.Error("Failed to parse request"))
			return
		}

		log.Info("Request body decoded", slog.Any("request", reqBody))

		if err := validator.New().Struct(reqBody); err != nil {
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)

			log.Error("Invalid request", sl.Err(err))
			render.JSON(writer, request, resp.ValidationError(validateErr))
			return
		}

		alias := reqBody.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
		}

		// ToDo: add condition for alias if is already exists in url table

		urlId, err := urlSaver.SaveUrl(reqBody.URL, reqBody.Alias)
		if err != nil {
			if errors.Is(err, repository.ErrURLExists) {
				log.Info("Url already exists", sl.Err(err))
				render.JSON(writer, request, resp.Error("Url already exists"))
				return
			}

			log.Error("Failed to add url", sl.Err(err))
			render.JSON(writer, request, resp.Error("Failed to add url"))
			return
		}

		log.Info("Url added", slog.Int64("UrlId", urlId))

		render.JSON(writer, request, Response{
			Response: resp.OK(),
			Alias:    alias,
		})
	}
}
