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

	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.staticDir)})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	logger.Info("starting server", slog.String("addr", cfg.addr))

	err := http.ListenAndServe(cfg.addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
