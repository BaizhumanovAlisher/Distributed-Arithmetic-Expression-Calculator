package handlers

import (
	"distributed_calculator/expression_manager"
	"distributed_calculator/expression_manager/agent"
	mwLogger "distributed_calculator/http_server/logger"
	"distributed_calculator/storage"
	"distributed_calculator/storage/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

func Routes(logger *slog.Logger, repo *postgresql.PostgresqlDB, redis *storage.RedisDB, manager *expression_manager.ExpressionManager, newAgent *agent.Agent) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.URLFormat)

	router.Post("/expression", createExpression(
		logger, repo.CreateExpression, redis.StoreIdempotencyToken, redis.RetrieveIdempotencyToken, manager))

	router.Get("/expressions", getExpressions(logger, repo.ReadExpressions))
	router.Get("/expressions/{id}", getExpression(logger, repo.ReadExpression))
	router.Get("/operations", getOperations(logger, repo.ReadOperations))
	router.Put("/operations", putOperations(logger, repo.UpdateOperation))
	router.Get("/mini-calculators", GetAllMiniCalculator(logger, newAgent.GetAllMiniCalculators))

	return router
}
