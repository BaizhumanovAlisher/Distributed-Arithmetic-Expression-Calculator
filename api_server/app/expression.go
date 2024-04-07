package app

import (
	"api_server/grpc_client"
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"internal/helpers"
	"internal/model"
	"internal/model/expression"
	"internal/validators"
	"log/slog"
	"net/http"
	"strconv"
)

type ExpressionReader interface {
	ReadExpressions(userId int64) ([]*expression.Expression, error)
	ReadExpression(id int) (*expression.Expression, error)
}

// createExpression data was reading in idempotencyExpressionPost and written to r.Context
func (app *Application) createExpression() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputExp, _ := r.Context().Value("expression").(string)
		err, _ := r.Context().Value("error").(error)

		if err != nil {
			app.log.Error("incorrect JSON file: ", err)
			apiError := model.NewAPIError("incorrect JSON file")

			rd := model.NewResponseData(
				http.StatusBadRequest,
				apiError,
			)

			app.updateContext(r, "response data", rd)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, apiError)

			return
		}

		err = validators.ValidateExpression(inputExp)
		if err != nil {
			app.log.Error("incorrect JSON file validating error: ", err)
			apiError := model.NewAPIError("incorrect JSON file, validating error")

			rd := model.NewResponseData(
				http.StatusBadRequest,
				apiError,
			)

			app.updateContext(r, "response data", rd)

			render.JSON(w, r, apiError)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		userId := int64(r.Context().Value(grpc_client.UserId).(float64))
		id, err := app.expressionSolver.CreateExpressionAndStartSolve(r.Context(), inputExp, userId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		idRespond := model.NewIdRespond(int64(id))
		rd := model.NewResponseData(
			http.StatusOK,
			idRespond,
		)

		app.updateContext(r, "response data", rd)

		render.JSON(w, r, idRespond)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func (app *Application) getExpressions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all expression")

		userId := int64(r.Context().Value(grpc_client.UserId).(float64))
		expressions, err := app.expressionReader.ReadExpressions(userId)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				app.log.Error("no expression", slog.Int64("userId", userId))
				w.WriteHeader(http.StatusNotFound)
				return
			}

			app.log.Error("internal server error", slog.String("err", err.Error()))
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError(err.Error()))
			return
		}

		app.log.Info("successful to get all expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, expressions)
		return
	}
}

func (app *Application) getExpression() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get expression")
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil || id < 1 {
			app.log.Info("id should be integer and bigger than 0")
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError("id should be integer"))
			return
		}

		exp, err := app.expressionReader.ReadExpression(id)

		if err != nil {
			if errors.Is(err, helpers.NoRowsErr) {
				app.log.Warn("no expression", slog.String("err", helpers.NoRowsErr.Error()))
				w.WriteHeader(http.StatusNotFound)
				return
			}

			app.log.Error("internal server error", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		userId := int64(r.Context().Value(grpc_client.UserId).(float64))
		if exp.UserId != userId {
			app.log.Warn("incorrect user id", slog.Int64("userId", userId), slog.Int64("expUserId", exp.UserId))
			w.WriteHeader(http.StatusForbidden)
			return
		}

		app.log.Info("successful to get expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, exp)
		return
	}
}
