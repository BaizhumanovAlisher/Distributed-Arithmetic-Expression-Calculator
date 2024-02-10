package model

type Operation struct {
	OperationKind    OperationType `json:"operationKind" validate:"required"`
	DurationInSecond int           `json:"durationInSecond" validate:"duration_in_sec"`
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

func IsAllowedOperation(operationType OperationType) bool {
	return operationType == Addition ||
		operationType == Subtraction ||
		operationType == Multiplication ||
		operationType == Division
}
