package expression

import "time"

type Expression struct {
	Id          int        `json:"id"`
	Expression  string     `json:"expression"`
	Answer      string     `json:"answer"`
	Status      Status     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt *time.Time `json:"completedAt"`
	UserId      int64      `json:"userId"`
}

type InputExpression struct {
	Expression string `json:"expression" validate:"required,expression"`
}

func NewExpressionInProcess(expression string, userId int64) *Expression {
	return &Expression{
		Expression: expression,
		Status:     InProcess,
		UserId:     userId,
	}
}

type Status string

const (
	Completed Status = "completed"
	InProcess        = "in process"
	Invalid          = "invalid"
)
