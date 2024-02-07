package model

type LeastExpression struct {
	Number1      float64       `json:"number1"`
	Number2      float64       `json:"number2"`
	Operation    OperationType `json:"operation"`
	IdExpression int           `json:"id_expression"`
}

func NewLeastExpression() *LeastExpression {
	return &LeastExpression{}
}
