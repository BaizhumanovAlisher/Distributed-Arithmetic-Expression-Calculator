package handlers

import (
	"api_server/expression_manager"
	"api_server/expression_manager/agent"
	mwLogger "api_server/http_server/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"internal/storage"
	"internal/storage/postgresql"
	"log/slog"
)

func Routes(logger *slog.Logger, repo *postgresql.PostgresqlDB, redis *storage.RedisDB, manager *expression_manager.ExpressionManager, newAgent *agent.Agent) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.URLFormat)

	router.Post("/expressions", idempotencyExpressionPost(
		createExpression(logger, manager, repo.CreateExpression),
		logger, redis))

	router.Get("/expressions", getExpressions(logger, repo.ReadExpressions))
	router.Get("/expressions/{id}", getExpression(logger, repo.ReadExpression))
	router.Get("/operations", getOperations(logger, repo.ReadOperations))
	router.Put("/operations", putOperations(logger, repo.UpdateOperation))
	router.Get("/mini-calculators", GetAllMiniCalculator(logger, newAgent.GetAllMiniCalculators))

	return router
}
