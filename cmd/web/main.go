package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"ed-henrique.snippetbox/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	config        config
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	cfg := getConfig()
	flag.Parse()

	loggerHandlerOptions := &slog.HandlerOptions{}

	if cfg.debug {
		loggerHandlerOptions.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, loggerHandlerOptions))

	db, err := openDB(*&cfg.dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config: cfg,
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		},
		templateCache: templateCache,
	}

	logger.Info("starting server", slog.String("addr", cfg.addr))

	err = http.ListenAndServe(cfg.addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
