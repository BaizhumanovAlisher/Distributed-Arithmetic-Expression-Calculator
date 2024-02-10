package main

import (
	"distributed_calculator/config"
	mwLogger "distributed_calculator/http_server/logger"
	"distributed_calculator/storage"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	logger := setupLogger()

	repo, err := storage.Postgresql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	//todo: run server

	operations, err := repo.ReadAllOperations()
	if err != nil {
		logger.Debug("%s", err)
	}

	for i := 0; i < len(operations); i++ {
		fmt.Println(operations[i])
	}
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
