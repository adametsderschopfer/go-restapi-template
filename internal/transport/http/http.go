package http

import (
	"app/internal/config"
	"log/slog"
	"net/http"
)

type Transport struct {
	Log      *slog.Logger
	Cfg      *config.Config
	Handlers http.Handler
}

func NewTransportServer(
	Log *slog.Logger,
	Cfg *config.Config,

	Handlers http.Handler,
) {
	t := Transport{
		Log:      Log,
		Cfg:      Cfg,
		Handlers: Handlers,
	}

	srv := &http.Server{
		Addr:         t.Cfg.Address,
		Handler:      t.Handlers,
		ReadTimeout:  t.Cfg.HTTPServer.Timeout,
		WriteTimeout: t.Cfg.HTTPServer.Timeout,
		IdleTimeout:  t.Cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		t.Log.Error("Failed to start server")
	}

	t.Log.Error("Server stopped")
}
