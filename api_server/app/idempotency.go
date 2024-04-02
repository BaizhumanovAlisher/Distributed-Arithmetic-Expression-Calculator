package app

import (
	"context"
	"github.com/go-chi/render"
	"internal/model"
	"internal/model/expression"
	"net/http"
)

func (app *Application) idempotencyExpressionPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		rd, err := app.redis.RetrieveIdempotencyToken(idempotencyToken, inputExpression.Expression)

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

			exp, ok := r.Context().Value("expression").(string)

			err := app.redis.StoreIdempotencyToken(idempotencyToken, exp, rd)
			if err != nil {
				app.log.Error("problem with redis: ", err)
			}
		}

		return
	}
}

func (app *Application) updateContext(r *http.Request, key string, value any) {
	ctx := r.Context()
	req := r.WithContext(context.WithValue(ctx, key, value))
	*r = *req
}
