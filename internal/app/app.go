package app

import (
	"app/internal/config"
	"app/internal/lib/logger/sl"
	"app/internal/repository/sqllite"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	Log *slog.Logger
	Cfg *config.Config

	Router http.Handler
}

func (app *App) New() {
	//repository := app.MustDatabaseLoad()

	app.Router = app.createRouter()
}

func (app *App) MustDatabaseLoad() *sqllite.Storage {
	storage, err := sqllite.New(app.Cfg.StoragePath)
	if err != nil {
		app.Log.Error("Failed to init repository", sl.Err(err))
		os.Exit(1)
	}
	return storage
}
