package app

import (
	"github.com/go-chi/render"
	"net/http"
)

func (app *Application) GetAllMiniCalculator() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.log.Info("start get all operations")

		miniCalculators, err := app.expressionSolver.GetCalculatorsStatus(r.Context())
		if err != nil {
			app.log.Warn("error getting calculators")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		app.log.Info("successful to get all operations")

		render.Status(r, http.StatusOK)
		render.JSON(w, r, miniCalculators)
	}
}
