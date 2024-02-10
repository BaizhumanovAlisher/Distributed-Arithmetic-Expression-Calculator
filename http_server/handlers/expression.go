package handlers

import (
	"database/sql"
	"distributed_calculator/http_server/validators"
	"distributed_calculator/model"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

func HandlerNewExpression(log *slog.Logger, expressionSaver func(expression *model.Expression) error, setResponseData func(idempotencyToken string, expression string, responseData *model.ResponseData) error, getResponseData func(idempotencyToken string, expression string) (*model.ResponseData, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputExpression model.InputExpression

		err := render.DecodeJSON(r.Body, &inputExpression)

		log.Info("request body decoded")

		idempotencyToken := r.Header.Get("X-Idempotency-Token")
		if checkCashedRespond(w, log, idempotencyToken, getResponseData, inputExpression) {
			return
		}

		if err != nil {
			apiError := model.NewAPIError("incorrect JSON file")
			log.Error("incorrect JSON file: %s", err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, apiError)

			bytes, _ := json.Marshal(apiError)
			cacheRespond(log, idempotencyToken, inputExpression.Expression, http.StatusInternalServerError, bytes, setResponseData)

			return
		}

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

			bytes, _ := json.Marshal(apiError)
			cacheRespond(log, idempotencyToken, inputExpression.Expression, http.StatusInternalServerError, bytes, setResponseData)

			return
		}

		//todo: add parser and start to solve

		render.Status(r, http.StatusOK)
		render.JSON(w, r, expression)

		bytes, _ := json.Marshal(expression)
		cacheRespond(log, idempotencyToken, inputExpression.Expression, http.StatusInternalServerError, bytes, setResponseData)

		log.Info("expression added", slog.Int("id", expression.Id))
	}
}

func checkCashedRespond(w http.ResponseWriter, log *slog.Logger, idempotencyToken string, getResponseData func(idempotencyToken string, expression string) (*model.ResponseData, error), inputExpression model.InputExpression) bool {
	log.Info("X-Idempotency-Token: %s", idempotencyToken)

	if idempotencyToken != "" {
		rd, err := getResponseData(idempotencyToken, inputExpression.Expression)

		if err != nil {
			log.Error("problem with redis: %s", err)
			return false
		}

		if rd != nil {
			log.Info("send cashed respond with status code: %d", rd.StatusCode)
			w.Write(rd.Body)
			w.WriteHeader(rd.StatusCode)
			return true
		}
	}
	return false
}

func cacheRespond(log *slog.Logger, idempotencyToken string, expression string, statusCode int, body []byte, setResponseData func(idempotencyToken string, expression string, responseData *model.ResponseData) error) {
	if idempotencyToken != "" {
		log.Info("start cache respond with %s and %s", idempotencyToken, expression)

		rd := model.NewResponseData(statusCode, body)
		err := setResponseData(idempotencyToken, expression, rd)
		if err != nil {
			log.Error("problem with redis: %s", err)
		}
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
