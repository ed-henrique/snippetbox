package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	cfg := getConfig()
	flag.Parse()

	loggerHandlerOptions := &slog.HandlerOptions{}

	if cfg.debug {
		loggerHandlerOptions.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, loggerHandlerOptions))

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", cfg.addr))

	err := http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
