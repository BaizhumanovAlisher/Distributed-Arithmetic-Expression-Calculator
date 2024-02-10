package handlers

import (
	"distributed_calculator/http_server/validators"
	"distributed_calculator/model"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type ExpressionSaveAndView interface {
	CreateExpression(*model.Expression) error
	ReadAllExpressions() ([]*model.Expression, error)
	ReadExpression(id int) (*model.Expression, error)
}

func New(log *slog.Logger, expressionSaver ExpressionSaveAndView) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputExpression model.InputExpression

		err := render.DecodeJSON(r.Body, &inputExpression)

		if err != nil {
			log.Error("incorrect JSON file: %s", err)
			render.Status(r, http.StatusBadRequest)
			return
		}

		log.Info("request body decoded")

		isCorrectValidated := validators.Validate(inputExpression.Expression)

		if !isCorrectValidated {
			expression := model.NewExpressionInvalid(inputExpression.Expression)
			err := expressionSaver.CreateExpression(expression)

			if err != nil {
				log.Error("%s", err)
			} else {
				log.Info("added expression to db: %+v", expression)
			}

			render.Status(r, http.StatusBadRequest)
		}

		//todo: add parser and start to solve

		expression := model.NewExpressionInProcess(inputExpression.Expression)
		render.Status(r, http.StatusOK)
		render.JSON(w, r, expression)

		log.Info("expression added", slog.Int("id", expression.Id))
	}
}
