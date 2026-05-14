package main

import (
	"log/slog"
	"os"
	"volodyadev/golang-restapi/internal/config"
	"volodyadev/golang-restapi/internal/lib/logger/sl"
	"volodyadev/golang-restapi/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Start app", slog.String("env", cfg.Env))
	log.Debug("Loglevel: DEBUG")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Storage initialization failed", sl.Err(err))
		os.Exit(1)
	}

	// Проверка создания url
	id, err := storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("Failed to save url", sl.Err(err))
	}
	log.Info("Saved url", slog.Int64("id", id))

	id, err = storage.SaveURL("https://google.com", "google")
	if err != nil {
		log.Error("Failed to save url", sl.Err(err))
	}

	// _ = storage
	// TODO: init logger: slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
