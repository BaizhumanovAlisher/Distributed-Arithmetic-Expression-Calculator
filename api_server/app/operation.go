package app

import (
	"github.com/go-chi/render"
	"internal/model"
	"internal/validators"
	"log/slog"
	"net/http"
)

func (app *Application) getOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all operations")

		operations, err := app.repo.ReadOperations()

		if err != nil {
			app.log.Error("error to get operations", slog.String("err", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		app.log.Info("successful to get all operations")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, operations)
	}
}

func (app *Application) putOperation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start put operation")

		var operation model.OperationWithDuration

		err := render.DecodeJSON(r.Body, &operation)

		if err != nil {
			app.log.Error("incorrect JSON file", slog.String("err", err.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError("incorrect JSON file"))
			return
		}

		app.log.Info("request body decoded")

		errValidating := validators.ValidateOperation(operation)

		if errValidating != nil {
			app.log.Error("err validating operation", slog.String("err", errValidating.Error()))
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError(errValidating.Error()))
			return
		}

		errDb := app.repo.UpdateOperation(&operation)

		if errDb != nil {
			app.log.Error("could not update operation", slog.String("operation", string(operation.OperationKind)))
			render.Status(r, http.StatusInternalServerError)
			return
		}

		app.log.Info("successful to update operation")
		w.WriteHeader(http.StatusOK)
	}
}
