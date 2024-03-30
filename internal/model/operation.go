package model

type OperationWithDuration struct {
	OperationKind    OperationType `json:"operationKind" validate:"required"`
	DurationInSecond int           `json:"durationInSecond" validate:"duration_in_sec"`
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

func DefineOperation(s rune) (OperationType, bool) {
	switch s {
	case '+':
		return Addition, true
	case '-':
		return Subtraction, true
	case '*':
		return Multiplication, true
	case '/':
		return Division, true

	default:
		return "", false
	}
}

// Precedence returns the precedence of an operation. Higher value means higher precedence.
func Precedence(op OperationType) int {
	switch op {
	case Addition, Subtraction:
		return 1
	case Multiplication, Division:
		return 2
	default:
		return 0
	}
}
