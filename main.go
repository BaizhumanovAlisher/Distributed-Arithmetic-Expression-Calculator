package main

import (
	"distributed_calculator/config"
	"distributed_calculator/http_server/handlers"
	mwLogger "distributed_calculator/http_server/logger"
	"distributed_calculator/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"log/slog"
	"net/http"
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

	setURLPatterns(router, logger, repo)

	logger.Info("start server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("failed to start")
	}
}

func setURLPatterns(router *chi.Mux, logger *slog.Logger, repo *storage.Sqlite) {
	router.Post("/expression", handlers.HandlerNewExpression(logger, repo.CreateExpression))
	router.Get("/expression", handlers.HandlerGetAllExpression(logger, repo.ReadAllExpressions))
	router.Get("/expression/{id}", handlers.HandlerGetExpression(logger, repo.ReadExpression))
	router.Get("/operation", handlers.HandlerGetAllOperations(logger, repo.ReadAllOperations))
	router.Put("/operation", handlers.HandlerPutOperations(logger, repo.UpdateOperation))
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
