package handlers

import (
	"database/sql"
	"distributed_calculator/expression_manager"
	"distributed_calculator/http_server/validators"
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func createExpression(log *slog.Logger, manager *expression_manager.ExpressionManager, expressionSaver func(expression *expression.Expression) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputExp, _ := r.Context().Value("expression").(string)
		err, _ := r.Context().Value("error").(error)

		if err != nil {
			log.Error("incorrect JSON file: ", err)
			apiError := model.NewAPIError("incorrect JSON file")

			rd := model.NewResponseData(
				http.StatusBadRequest,
				apiError,
			)

			updateContext(r, "response data", rd)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, apiError)

			return
		}

		err = validators.ValidateExpression(inputExp)
		if err != nil {
			log.Error("incorrect JSON file validating error: ", err)
			apiError := model.NewAPIError("incorrect JSON file, validating error")

			rd := model.NewResponseData(
				http.StatusBadRequest,
				apiError,
			)

			updateContext(r, "response data", rd)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, apiError)

			return
		}

		expressionFull := expression.NewExpression(inputExp)
		err = expressionSaver(expressionFull)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return
		}

		manager.StartSolveConcurrently(expressionFull)

		rd := model.NewResponseData(
			http.StatusOK,
			expressionFull,
		)

		updateContext(r, "response data", rd)

		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, expressionFull)
		return
	}
}

func getExpressions(log *slog.Logger, expressionReader func() ([]*expression.Expression, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start get all expression")

		expressions, err := expressionReader()

		if errors.Is(err, sql.ErrNoRows) {
			log.Error("error to get expression: %s", err)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, model.NewAPIError("no expressions"))
			return
		}

		log.Info("successful to get all expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, expressions)
	}
}

func getExpression(log *slog.Logger, expressionReader func(int) (*expression.Expression, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start get expression")
		idStr := chi.URLParam(r, "id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			log.Error("id should be integer and bigger than 0")
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError("id should be integer"))
			return
		}

		exp, err := expressionReader(id)

		if errors.Is(err, sql.ErrNoRows) {
			log.Error("error to get expression: %s", err)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, model.NewAPIError("no expression with this id"))
			return
		}

		log.Info("successful to get expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, exp)
	}
}
