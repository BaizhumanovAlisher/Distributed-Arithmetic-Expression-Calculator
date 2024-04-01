package validators

import (
	"errors"
)

func ValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("empty username")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) == 0 {
		return errors.New("empty password")
	}

	return nil
}
