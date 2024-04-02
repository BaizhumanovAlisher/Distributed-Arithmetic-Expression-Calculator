package app

import (
	"database/sql"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"internal/helpers"
	model2 "internal/model"
	"internal/model/expression"
	"internal/validators"
	"net/http"
	"strconv"
)

func (app *Application) createExpression() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputExp, _ := r.Context().Value("expression").(string)
		err, _ := r.Context().Value("error").(error)

		if err != nil {
			app.log.Error("incorrect JSON file: ", err)
			apiError := helpers.NewAPIError("incorrect JSON file")

			rd := model2.NewResponseData(
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
			apiError := helpers.NewAPIError("incorrect JSON file, validating error")

			rd := model2.NewResponseData(
				http.StatusBadRequest,
				apiError,
			)

			app.updateContext(r, "response data", rd)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, apiError)

			return
		}

		expressionFull := expression.NewExpressionInProcess(inputExp)
		_, err = app.repo.CreateExpression(expressionFull)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return
		}

		app.manager.StartSolveConcurrently(expressionFull)

		rd := model2.NewResponseData(
			http.StatusOK,
			expressionFull,
		)

		app.updateContext(r, "response data", rd)

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, expressionFull)
		return
	}
}

func (app *Application) getExpressions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all expression")

		// todo: add reading from context user_id
		var userId int64
		expressions, err := app.repo.ReadExpressions(userId)

		if errors.Is(err, sql.ErrNoRows) {
			app.log.Error("error to get expression: %s", err)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, helpers.NewAPIError("no expressions"))
			return
		}

		app.log.Info("successful to get all expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, expressions)
	}
}

func (app *Application) getExpression() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get expression")
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			app.log.Error("id should be integer and bigger than 0")
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, helpers.NewAPIError("id should be integer"))
			return
		}

		// todo: check user_id to validate access
		exp, err := app.repo.ReadExpression(id)

		if errors.Is(err, sql.ErrNoRows) {
			app.log.Error("error to get expression: %s", err)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, helpers.NewAPIError("no expression with this id"))
			return
		}

		app.log.Info("successful to get expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, exp)
	}
}
