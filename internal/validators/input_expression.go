package validators

import (
	"errors"
	"unicode"
)

func ValidateExpression(expression string) error {
	if len(expression) == 0 {
		return errors.New("empty expression")
	}

	err := notContainExtraCharacters(expression)

	if err != nil {
		return err
	}

	err = bracketsAreCorrect(expression)
	return err
}

func notContainExtraCharacters(exp string) error {
	for _, char := range exp {
		if !(unicode.IsDigit(char) || isAllowedChar(char)) {
			return errors.New("expression should not contain extra symbols")
		}
	}

	return nil
}

// AllowedCharacters - the array is sorted
var AllowedCharacters = []rune{'(', ')', '*', '+', '-', '.', '/'}

func isAllowedChar(c rune) bool {
	for _, char := range AllowedCharacters {
		if c == char {
			return true
		}
	}

	return false
}

func bracketsAreCorrect(exp string) error {
	var stack []rune

	for _, r := range exp {
		switch r {
		case '(':
			stack = append(stack, r)
		case ')':
			if len(stack) == 0 {
				return errors.New("brackets should be correct")
			}

			last := stack[len(stack)-1]
			if r == ')' && last != '(' {
				return errors.New("brackets should be correct")
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return errors.New("brackets should be correct")
	}

	return nil
}
