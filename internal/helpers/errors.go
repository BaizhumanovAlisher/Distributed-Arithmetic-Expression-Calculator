package helpers

import (
	"fmt"
)

var (
	InvalidCredentialsErr      = fmt.Errorf("invalid credentials")
	UsernameExistErr           = fmt.Errorf("user name exist already")
	NoRowsErr                  = fmt.Errorf("not found")
	InternalErr                = fmt.Errorf("internal error")
	InvalidArgumentUserNameErr = fmt.Errorf("length of name should be longer than 7")
	InvalidArgumentPasswordErr = fmt.Errorf("length of name should be longer than 7")
)
