package app

import (
	"log/slog"
	"net/http"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqllite"
)

type App struct {
	Log *slog.Logger
	Cfg *config.Config

	Router http.Handler
}

func (app *App) New() {
	//storage := app.MustDatabaseLoad()

	app.Router = app.createRouter()
}

func (app *App) MustDatabaseLoad() *sqllite.Storage {
	storage, err := sqllite.New(app.Cfg.StoragePath)
	if err != nil {
		app.Log.Error("Failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	return storage
}
