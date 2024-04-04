package validators

import (
	"internal/helpers"
)

func ValidateUsername(username string) error {
	if len([]rune(username)) < 8 {
		return helpers.InvalidArgumentUserNameErr
	}

	return nil
}

func ValidatePassword(password string) error {
	if len([]rune(password)) < 8 {
		return helpers.InvalidArgumentPasswordErr
	}

	return nil
}
