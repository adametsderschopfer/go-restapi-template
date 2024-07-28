package app

import (
	"app/internal/config"
	"app/internal/lib/logger/sl"
	"app/internal/repository/postgresql"
	"fmt"
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
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		app.Cfg.Postgres.Host,
		app.Cfg.Postgres.Port,
		app.Cfg.Postgres.User,
		app.Cfg.Postgres.Password,
		app.Cfg.Postgres.DBName,
	)

	repository, err := postgresql.New(connectionString)
	if err != nil {
		app.Log.Error("Failed to init repository", sl.Err(err))
		os.Exit(1)
	}

	app.Repository = repository
	app.Router = app.createRouter()
}
