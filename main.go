package main

import (
	"distributed_calculator/config"
	"distributed_calculator/expression_manager"
	"distributed_calculator/expression_manager/agent"
	"distributed_calculator/http_server/handlers"
	"distributed_calculator/storage"
	"distributed_calculator/storage/postgresql"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()
	logger := NewLogger()

	repo, err := postgresql.Postgresql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := storage.Redis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	newAgent := agent.NewAgent(cfg.Agent.CountOperation)
	expressionManager, err := expression_manager.NewExpressionManager(newAgent, repo)

	if err != nil {
		log.Fatal(err)
	}

	router := handlers.Routes(logger, repo, redis, expressionManager, newAgent)

	logger.Info("start server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func NewLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
