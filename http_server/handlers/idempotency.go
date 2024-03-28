package handlers

import (
	"context"
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"distributed_calculator/storage"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func idempotencyExpressionPost(next http.HandlerFunc, logger *slog.Logger, redis *storage.RedisDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idempotencyToken := r.Header.Get("X-Idempotency-Token")

		var inputExpression expression.InputExpression
		err := render.DecodeJSON(r.Body, &inputExpression)

		updateContext(r, "expression", inputExpression.Expression)
		updateContext(r, "error", err)

		if idempotencyToken == "" {
			next.ServeHTTP(w, r)
			return
		}
		logger.Debug("idempotency token: ", idempotencyToken)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		rd, err := redis.RetrieveIdempotencyToken(idempotencyToken, inputExpression.Expression)

		if err != nil {
			logger.Error("problem with redis: ", err)
			next.ServeHTTP(w, r)
			return
		}

		if rd != nil {
			logger.Info("send cashed respond")

			w.WriteHeader(rd.StatusCode)
			render.JSON(w, r, rd.Body)
		} else {
			next.ServeHTTP(w, r)

			rd, ok := r.Context().Value("response data").(*model.ResponseData)
			if !ok {
				logger.Error("No saved response data")
				return
			}

			exp, ok := r.Context().Value("expression").(string)

			err := redis.StoreIdempotencyToken(idempotencyToken, exp, rd)
			if err != nil {
				logger.Error("problem with redis: ", err)
			}
		}

		return
	}
}

func updateContext(r *http.Request, key string, value any) {
	ctx := r.Context()
	req := r.WithContext(context.WithValue(ctx, key, value))
	*r = *req
}
