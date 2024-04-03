package validators

import (
	"internal/helpers"
)

func ValidateUsername(username string) error {
	if len([]rune(username)) < 8 {
		return helpers.InvalidArgumentUserName
	}

	return nil
}

func ValidatePassword(password string) error {
	if len([]rune(password)) < 8 {
		return helpers.InvalidArgumentPassword
	}

	return nil
}
