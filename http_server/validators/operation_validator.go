package validators

import (
	"distributed_calculator/model"
	"errors"
)

func ValidateOperation(operation model.Operation) error {
	if operation.DurationInSecond < 0 {
		return errors.New("operation duration should be bigger than 0")
	}
	return nil
}
