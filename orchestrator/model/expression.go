package model

import "time"

type Expression struct {
	Id          int       `json:"id"`
	Expression  string    `json:"expression"`
	Answer      string    `json:"answer"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

func NewExpression() *Expression {
	return &Expression{}
}

type Status string

const (
	Completed Status = "completed"
	InProcess        = "in process"
	Invalid          = "invalid"
)
