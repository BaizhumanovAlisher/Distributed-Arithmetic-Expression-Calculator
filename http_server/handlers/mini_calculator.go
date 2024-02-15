package handlers

import (
	"distributed_calculator/model/expression"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

func HandlerGetAllMiniCalculator(log *slog.Logger, miniCalculatorReader func() []*expression.MiniCalculator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("start get all operations")

		miniCalculators := miniCalculatorReader()

		log.Info("successful to get all operations")
		render.Status(r, http.StatusOK)
		render.JSON(w, r, miniCalculators)
	}
}
