package app

import (
	"api_server/expression_manager"
	"api_server/expression_manager/agent"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	mwLogger "internal/helpers"
	"internal/storage"
	"internal/storage/postgresql"
	"log/slog"
)

type Application struct {
	log      *slog.Logger
	repo     *postgresql.PostgresqlDB
	redis    *storage.RedisDB
	manager  *expression_manager.ExpressionManager
	newAgent *agent.Agent
}

func NewApplication(log *slog.Logger, repo *postgresql.PostgresqlDB, redis *storage.RedisDB, manager *expression_manager.ExpressionManager, newAgent *agent.Agent) *Application {
	return &Application{
		log:      log,
		repo:     repo,
		redis:    redis,
		manager:  manager,
		newAgent: newAgent,
	}
}

func (app *Application) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(mwLogger.NewLoggerMiddleware(app.log))
	router.Use(middleware.URLFormat)

	router.Post("/expressions", app.idempotencyExpressionPost(app.createExpression()))

	router.Get("/expressions", app.getExpressions())
	router.Get("/expressions/{id}", app.getExpression())
	router.Get("/operations", app.getOperations())
	router.Put("/operations", app.putOperations())
	router.Get("/mini-calculators", app.GetAllMiniCalculator())

	//todo: add sing_in and login
	return router
}
