package validators

import (
	"distributed_calculator/model"
	"errors"
)

func ValidateOperation(operation model.Operation) error {
	if operation.DurationInSecond < 0 || operation.DurationInSecond > 30 {
		return errors.New("operation duration should be more than 0 and less 30")
	}

	if !model.IsAllowedOperation(operation.OperationKind) {
		return errors.New("it is not allowed operation")
	}

	return nil
}
