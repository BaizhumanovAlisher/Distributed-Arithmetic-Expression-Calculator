package validators

import (
	"errors"
	"internal/model"
)

func ValidateOperation(operation model.OperationWithDuration) error {
	if operation.DurationInSecond < 0 {
		return errors.New("operation duration should be more than 0")
	}

	if !model.IsAllowedOperation(operation.OperationKind) {
		return errors.New("it is not allowed operation")
	}

	return nil
}
