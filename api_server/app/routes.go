package app

import (
	"api_server/expression_manager"
	"api_server/expression_manager/agent"
	"api_server/grpc_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"internal/storage"
	"internal/storage/postgresql"
	"log/slog"
)

type Application struct {
	log         *slog.Logger
	repo        *postgresql.PostgresqlDB
	redis       *storage.RedisDB
	manager     *expression_manager.ExpressionManager
	newAgent    *agent.Agent
	authService *grpc_client.AuthService
}

func NewApplication(
	log *slog.Logger, repo *postgresql.PostgresqlDB, redis *storage.RedisDB,
	manager *expression_manager.ExpressionManager, newAgent *agent.Agent,
	authService *grpc_client.AuthService) *Application {

	return &Application{
		log:         log,
		repo:        repo,
		redis:       redis,
		manager:     manager,
		newAgent:    newAgent,
		authService: authService,
	}
}

func (app *Application) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(NewLoggerMiddleware(app.log))
	router.Use(middleware.URLFormat)

	router.Post("/expressions", app.idempotencyExpressionPost(app.createExpression()))

	router.Get("/expressions", app.getExpressions())
	router.Get("/expressions/{id}", app.getExpression())
	router.Get("/operations", app.getOperations())
	router.Put("/operations", app.putOperation())
	router.Get("/mini_calculators", app.GetAllMiniCalculator())

	//todo: add sing_in and login
	return router
}
