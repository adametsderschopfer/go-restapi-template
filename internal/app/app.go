package app

import (
	"app/internal/config"
	"app/internal/lib/logger/sl"
	"app/internal/repository/postgresql"
	"log/slog"
	"net/http"
	"os"
)

type App struct {
	Log *slog.Logger
	Cfg *config.Config

	Router     http.Handler
	Repository *postgresql.Repository
}

func (app *App) New() {
	repository, err := postgresql.New(postgresql.CreateConnectionString(app.Cfg))
	if err != nil {
		app.Log.Error("Failed to init repository", sl.Err(err))
		os.Exit(1)
	}

	app.Repository = repository
	app.Router = app.createRouter()
}
