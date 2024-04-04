package app

import (
	"errors"
	"github.com/go-chi/render"
	"internal/helpers"
	"internal/model"
	"net/http"
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
			if errors.Is(err, helpers.UsernameExistErr) || errors.Is(err, helpers.InvalidArgumentUserNameErr) || errors.Is(err, helpers.InvalidArgumentUserNameErr) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, model.NewAPIError(helpers.InternalErr.Error()))
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
				render.JSON(w, r, model.NewAPIError(helpers.InternalErr.Error()))
				return
			}

			if errors.Is(err, helpers.InvalidArgumentUserNameErr) || errors.Is(err, helpers.InvalidArgumentUserNameErr) {
				render.Status(r, http.StatusBadRequest)
				render.JSON(w, r, model.NewAPIError(helpers.InternalErr.Error()))
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
