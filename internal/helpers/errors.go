package helpers

import "fmt"

var (
	InvalidCredentialsErr   = fmt.Errorf("invalid credentials")
	UsernameExistErr        = fmt.Errorf("user name exist already")
	NoRowsErr               = fmt.Errorf("not found")
	InternalErr             = fmt.Errorf("internal error")
	InvalidArgumentUserName = fmt.Errorf("length of name should be longer than 7")
	InvalidArgumentPassword = fmt.Errorf("length of name should be longer than 7")
)

type APIError struct {
	ApiErr string `json:"apiError"`
	Id     *int   `json:"id,omitempty"`
}

func NewAPIError(s string) *APIError {
	return &APIError{ApiErr: s}
}
