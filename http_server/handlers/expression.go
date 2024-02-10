package handlers

import (
	"database/sql"
	"distributed_calculator/http_server/validators"
	"distributed_calculator/model"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func HandlerNewExpression(log *slog.Logger, expressionSaver func(expression *model.Expression) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputExpression model.InputExpression

		//todo: add token idempotent X-Idempotency-Token

		err := render.DecodeJSON(r.Body, &inputExpression)

		if err != nil {
			log.Error("incorrect JSON file: %s", err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"description": "incorrect JSON file"})
			return
		}

		log.Info("request body decoded")

		errValidating := validators.ValidateExpression(inputExpression.Expression)

		var expression *model.Expression

		if errValidating != nil {
			expression = model.NewExpressionInvalid(inputExpression.Expression)
		} else {
			expression = model.NewExpressionInProcess(inputExpression.Expression)
		}

		errDb := expressionSaver(expression)

		if errDb != nil {
			log.Error("%s", errDb)

			apiErr := model.NewAPIError("problem with database")
			apiErr.Id = &expression.Id

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, apiErr)
			return
		} else {
			log.Info("added expression to db: %+v", expression)
		}

		if errValidating != nil {
			apiError := model.NewAPIError(errValidating.Error())
			apiError.Id = &expression.Id

			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, apiError)
			return
		}

		//todo: add parser and start to solve

		render.Status(r, http.StatusOK)
		render.JSON(w, r, expression)

		log.Info("expression added", slog.Int("id", expression.Id))
	}
}

func HandlerGetAllExpression(log *slog.Logger, expressionReader func() ([]*model.Expression, error)) http.HandlerFunc {
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

func HandlerGetExpression(log *slog.Logger, expressionReader func(int) (*model.Expression, error)) http.HandlerFunc {
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

		expression, err := expressionReader(id)

		if errors.Is(err, sql.ErrNoRows) {
			log.Error("error to get expression: %s", err)
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, model.NewAPIError("no expression with this id"))
			return
		}

		log.Info("successful to get expressions")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, expression)
	}
}
