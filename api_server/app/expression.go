package app

import (
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
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, apiError)

			return
		}

		expressionFull := expression.NewExpressionInProcess(inputExp)
		id, err := app.repo.CreateExpression(expressionFull)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return
		}

		app.manager.StartSolveConcurrently(expressionFull)

		idRespond := model.NewIdRespond(int64(id))
		rd := model.NewResponseData(
			http.StatusOK,
			idRespond,
		)

		app.updateContext(r, "response data", rd)

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, idRespond)
		return
	}
}

func (app *Application) getExpressions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all expression")

		// todo: add reading from context user_id
		var userId int64
		expressions, err := app.repo.ReadExpressions(userId)

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

		// todo: check user_id to validate access. permission may be denied
		exp, err := app.repo.ReadExpression(id)

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

		app.log.Info("successful to get expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, exp)
		return
	}
}
