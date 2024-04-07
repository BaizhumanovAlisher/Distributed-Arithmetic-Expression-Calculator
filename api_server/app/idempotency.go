package app

import (
	"api_server/grpc_client"
	"context"
	"github.com/go-chi/render"
	"internal/model"
	"internal/model/expression"
	"net/http"
)

type IdempotencyTokenRepo interface {
	StoreIdempotencyToken(generatedKey string, rd *model.ResponseData) error
	RetrieveIdempotencyToken(generatedKey string) (*model.ResponseData, error)
	GenerateTokenKey(idempotencyToken string, expression string, userId int64) string
}

func (app *Application) idempotencyExpressionPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idempotencyToken := r.Header.Get("X-Idempotency-Token")

		var inputExpression expression.InputExpression
		err := render.DecodeJSON(r.Body, &inputExpression)

		app.updateContext(r, "expression", inputExpression.Expression)
		app.updateContext(r, "error", err)

		if idempotencyToken == "" {
			next.ServeHTTP(w, r)
			return
		}
		app.log.Debug("idempotency token: ", idempotencyToken)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		userId := int64(r.Context().Value(grpc_client.UserId).(float64))
		generatedKey := app.idempotencyTokenRepo.GenerateTokenKey(idempotencyToken, inputExpression.Expression, userId)

		rd, err := app.idempotencyTokenRepo.RetrieveIdempotencyToken(generatedKey)

		if err != nil {
			app.log.Error("problem with redis: ", err)
			next.ServeHTTP(w, r)
			return
		}

		if rd != nil {
			app.log.Info("send cashed respond")

			w.WriteHeader(rd.StatusCode)
			render.JSON(w, r, rd.Body)
		} else {
			next.ServeHTTP(w, r)

			rd, ok := r.Context().Value("response data").(*model.ResponseData)
			if !ok {
				app.log.Error("No saved response data")
				return
			}

			userId := int64(r.Context().Value(grpc_client.UserId).(float64))
			generatedKey := app.idempotencyTokenRepo.GenerateTokenKey(idempotencyToken, inputExpression.Expression, userId)

			err := app.idempotencyTokenRepo.StoreIdempotencyToken(generatedKey, rd)
			if err != nil {
				app.log.Error("problem with redis: ", err)
			}
		}

		return
	})
}

func (app *Application) updateContext(r *http.Request, key string, value any) {
	ctx := r.Context()
	req := r.WithContext(context.WithValue(ctx, key, value))
	*r = *req
}
