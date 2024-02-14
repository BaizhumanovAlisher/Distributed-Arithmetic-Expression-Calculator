package handlers

import (
	"database/sql"
	"distributed_calculator/http_server/validators"
	"distributed_calculator/model"
	"errors"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func HandlerGetAllOperations(log *slog.Logger, operationReader func() ([]*model.OperationWithDuration, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start get all operations")

		operations, err := operationReader()

		if errors.Is(err, sql.ErrNoRows) {
			log.Error("error to get operations: %s", err)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError("no operations"))
			return
		}

		log.Info("successful to get all operations")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, operations)
	}
}

func HandlerPutOperations(log *slog.Logger, operationUpdate func(operation *model.OperationWithDuration) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start put operations")

		var operation model.OperationWithDuration

		err := render.DecodeJSON(r.Body, &operation)

		if err != nil {
			log.Error("incorrect JSON file: %s", err)
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, model.NewAPIError("incorrect JSON file"))
			return
		}

		log.Info("request body decoded")

		errValidating := validators.ValidateOperation(operation)

		if errValidating != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError(errValidating.Error()))
			return
		}

		errDb := operationUpdate(&operation)

		if errDb != nil {
			log.Error("could not update operation: %+v", operation)
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, model.NewAPIError("could not update operation"))
		}

		log.Info("successful to update operation")
		w.WriteHeader(http.StatusOK)
	}
}
