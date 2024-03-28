package expression

import (
	"distributed_calculator/internal/model"
)

type LeastExpression struct {
	Number1          float64             `json:"number1"`
	Number2          float64             `json:"number2"`
	Operation        model.OperationType `json:"operation"`
	IdExpression     int                 `json:"idExpression"`
	DurationInSecond int                 `json:"durationInSecond"`
	Result           float64             `json:"-"`
	ResultIsCorrect  chan bool           `json:"-"`
}

func NewLeastExpression(number1 float64, number2 float64, operation model.OperationType, idExpression int, durationInSecond int) *LeastExpression {
	return &LeastExpression{
		Number1:          number1,
		Number2:          number2,
		Operation:        operation,
		IdExpression:     idExpression,
		DurationInSecond: durationInSecond,
		ResultIsCorrect:  make(chan bool),
	}
}
