package main

import (
	"distributed_calculator/config"
	"distributed_calculator/expression_manager"
	"distributed_calculator/expression_manager/agent"
	"distributed_calculator/http_server/handlers"
	mwLogger "distributed_calculator/http_server/logger"
	"distributed_calculator/storage"
	"distributed_calculator/storage/postgreql"
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

	repo, err := postgreql.Postgresql(cfg)
	if err != nil {
		log.Fatal(err)
	}

	redis, err := storage.Redis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	newAgent := agent.NewAgent(cfg.Agent.CountOperation)
	expressionManager, err := expression_manager.NewExpressionManager(
		newAgent, repo.ReadOperation, repo.UpdateExpression, repo.ReadAllExpressionsWithStatus)

	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	setURLPatterns(router, logger, repo, redis, expressionManager, newAgent)

	logger.Info("start server", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:        cfg.HTTPServer.Address,
		Handler:     router,
		IdleTimeout: cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("failed to start")
	}
}

func setURLPatterns(router *chi.Mux, logger *slog.Logger, repo *postgreql.PostgresqlDB, redis *storage.RedisDB, manager *expression_manager.ExpressionManager, newAgent *agent.Agent) {

	router.Post("/expression", handlers.HandlerNewExpression(
		logger, repo.CreateExpression, redis.StoreIdempotencyToken, redis.RetrieveIdempotencyToken, manager))

	router.Get("/expression", handlers.HandlerGetAllExpression(logger, repo.ReadAllExpressions))
	router.Get("/expression/{id}", handlers.HandlerGetExpression(logger, repo.ReadExpression))
	router.Get("/operation", handlers.HandlerGetAllOperations(logger, repo.ReadAllOperations))
	router.Put("/operation", handlers.HandlerPutOperations(logger, repo.UpdateOperation))
	router.Get("/mini-calculator", handlers.HandlerGetAllMiniCalculator(logger, newAgent.GetAllMiniCalculators))
}

func setupLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{Level: slog.LevelDebug}))
}
