package model

type Operation struct {
	OperationKind    OperationType `json:"operationKind"`
	DurationInSecond int           `json:"durationInSecond"`
}

func NewOperation() *Operation {
	return &Operation{}
}

type OperationType string

const (
	Addition       OperationType = "addition"
	Subtraction                  = "subtraction"
	Multiplication               = "multiplication"
	Division                     = "division"
)
