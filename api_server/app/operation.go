package app

import (
	"database/sql"
	"errors"
	"github.com/go-chi/render"
	"internal/helpers"
	model2 "internal/model"
	"internal/validators"
	"net/http"
)

func (app *Application) getOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all operations")

		operations, err := app.repo.ReadOperations()

		if errors.Is(err, sql.ErrNoRows) {
			app.log.Error("error to get operations: %s", err)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, helpers.NewAPIError("no operations"))
			return
		}

		app.log.Info("successful to get all operations")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, operations)
	}
}

func (app *Application) putOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start put operations")

		var operation model2.OperationWithDuration

		err := render.DecodeJSON(r.Body, &operation)

		if err != nil {
			app.log.Error("incorrect JSON file: %s", err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, helpers.NewAPIError("incorrect JSON file"))
			return
		}

		app.log.Info("request body decoded")

		errValidating := validators.ValidateOperation(operation)

		if errValidating != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, helpers.NewAPIError(errValidating.Error()))
			return
		}

		errDb := app.repo.UpdateOperation(&operation)

		if errDb != nil {
			app.log.Error("could not update operation: %+v", operation)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, helpers.NewAPIError("could not update operation"))
		}

		app.log.Info("successful to update operation")
		w.WriteHeader(http.StatusOK)
	}
}
