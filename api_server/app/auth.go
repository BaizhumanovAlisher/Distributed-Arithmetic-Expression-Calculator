package app

import (
	"errors"
	"github.com/go-chi/render"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"internal/helpers"
	"internal/model"
	"log/slog"
	"net/http"
	"time"
)

func (app *Application) registerUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserCredentials

		if err := render.DecodeJSON(r.Body, &user); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError(err.Error()))
			return
		}

		id, err := app.authService.Register(r.Context(), user.Name, user.Password)
		if err != nil {
			if errors.Is(err, helpers.UsernameExistErr) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, model.NewAPIError(helpers.UsernameExistErr.Error()))
				return
			}

			if errors.Is(err, helpers.InvalidArgumentUserNameErr) || errors.Is(err, helpers.InvalidArgumentPasswordErr) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, model.NewAPIError(err.Error()))
				return
			}

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError(helpers.InternalErr.Error()))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewIdRespond(id))
		return
	}
}

func (app *Application) generateJWT() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.UserCredentials

		if err := render.DecodeJSON(r.Body, &user); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError(err.Error()))
			return
		}

		token, err := app.authService.Login(r.Context(), user.Name, user.Password)
		if err != nil {
			if errors.Is(err, helpers.InvalidCredentialsErr) {
				render.Status(r, http.StatusForbidden)
				render.JSON(w, r, model.NewAPIError(helpers.InvalidCredentialsErr.Error()))
				return
			}

			if errors.Is(err, helpers.InvalidArgumentUserNameErr) || errors.Is(err, helpers.InvalidArgumentPasswordErr) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, model.NewAPIError(err.Error()))
				return
			}

			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError(helpers.InternalErr.Error()))
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, model.NewTokenRespond(token))
		return
	}
}

func (app *Application) middlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bearerExtractor request.BearerExtractor
		tokenString, err := bearerExtractor.ExtractToken(r)

		if err != nil || len(tokenString) == 0 {
			slog.Warn("Error extracting token", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		app.log.Info("successfully getting token", slog.String("token", tokenString))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return app.authService.Secret, nil
		})

		if err != nil {
			slog.Warn("Error extracting token", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			slog.Warn("Invalid token", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		values := token.Header
		timeExpiry, ok := values["exp"].(time.Time)
		if !ok || timeExpiry.After(time.Now()) {
			slog.Warn("Expired token or does not have \"exp\"", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userId, ok := values["userId"].(int64)
		if !ok || userId == 0 {
			slog.Error("do not have user id", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		app.updateContext(r, "userId", userId)
		next.ServeHTTP(w, r)
		return
	})
}
