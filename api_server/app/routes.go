package app

import (
	"api_server/grpc_client"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/alice"
	"internal/storage"
	"internal/storage/postgresql"
	"log/slog"
	"net/http"
)

type Application struct {
	log              *slog.Logger
	repo             *postgresql.PostgresqlDB
	redis            *storage.RedisDB
	authService      *grpc_client.AuthService
	expressionSolver *grpc_client.ExpressionSolver
}

func NewApplication(log *slog.Logger, repo *postgresql.PostgresqlDB, redis *storage.RedisDB, authService *grpc_client.AuthService, solver *grpc_client.ExpressionSolver) *Application {

	return &Application{
		log:              log,
		repo:             repo,
		redis:            redis,
		authService:      authService,
		expressionSolver: solver,
	}
}

func (app *Application) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(NewLoggerMiddleware(app.log))
	router.Use(middleware.URLFormat)

	router.Post("/sign_up", app.registerUser())
	router.Post("/login", app.generateJWT())

	protected := alice.New(app.middlewareAuth)

	router.Method(http.MethodPost, "/expressions", protected.Then(alice.New(app.idempotencyExpressionPost).ThenFunc(app.createExpression())))
	router.Method(http.MethodGet, "/expressions", protected.Then(app.getExpressions()))
	router.Method(http.MethodGet, "/expressions/{id}", protected.Then(app.getExpression()))
	router.Method(http.MethodGet, "/operations", protected.Then(app.getOperations()))
	router.Method(http.MethodPut, "/operations", protected.Then(app.putOperation()))
	router.Method(http.MethodGet, "/mini_calculators", protected.Then(app.GetAllMiniCalculator()))

	return router
}
