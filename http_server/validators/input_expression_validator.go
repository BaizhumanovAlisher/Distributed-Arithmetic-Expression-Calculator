package validators

import (
	"context"
	"errors"
	"log/slog"
	"strings"
	"sync"
	"unicode"
)

func Validate(expression string) bool {
	slog.Info("start validating expression: %s", expression)
	if len(expression) == 0 {
		slog.Info("empty expression")
		return false
	}

	var wg sync.WaitGroup
	errorCh := make(chan error)
	defer close(errorCh)
	ctx := context.Background()
	exp := []rune(expression)

	wg.Add(6)

	concurrentErrorDetection(startWithNumber, exp, &wg, errorCh, ctx)
	concurrentErrorDetection(notContainExtraCharacters, exp, &wg, errorCh, ctx)
	concurrentErrorDetection(digitStartWithZero, exp, &wg, errorCh, ctx)
	concurrentErrorDetection(wholeExpressionIsDigit, exp, &wg, errorCh, ctx)
	concurrentErrorDetection(divideByZero, exp, &wg, errorCh, ctx)
	concurrentErrorDetection(bracketsAreCorrect, exp, &wg, errorCh, ctx)

	wg.Wait()

	if len(errorCh) == 0 {
		slog.Info("successful validation expression: %s", expression)
		return true
	} else {
		var descriptionErrors strings.Builder

		descriptionErrors.WriteString((<-errorCh).Error())

		for err := range errorCh {
			descriptionErrors.WriteString(",")
			descriptionErrors.WriteString(err.Error())
		}

		slog.Info("unsuccessful validation: %s - [%s]", expression, descriptionErrors.String())
		return false
	}
}

type CheckFunc func([]rune) error

func concurrentErrorDetection(checkExpressionToError CheckFunc, exp []rune, wg *sync.WaitGroup, errorCh chan error, ctx context.Context) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	default:
		err := checkExpressionToError(exp)

		if err != nil {
			errorCh <- err

			_, cancel := context.WithCancel(ctx)
			cancel()
			return
		}
	}
}

func startWithNumber(exp []rune) error {
	firstChar := exp[0]
	if !unicode.IsDigit(firstChar) {
		return errors.New("expression should start with number")
	}

	return nil
}

func notContainExtraCharacters(exp []rune) error {
	for _, char := range exp {
		if !(unicode.IsDigit(char) || isAllowedChar(char)) {
			return errors.New("expression should not contain extra symbols")
		}
	}

	return nil
}

// AllowedCharacters - the array is sorted
var AllowedCharacters = []rune{'(', ')', '*', '+', '-', '.', '/'}

// todo: do binary search
func isAllowedChar(c rune) bool {
	for _, char := range AllowedCharacters {
		if c == char {
			return true
		}
	}

	return false
}

func digitStartWithZero(exp []rune) error {
	if exp[0] == '0' && unicode.IsDigit(exp[1]) {
		return errors.New("digit should not start with zero")
	}

	for i := 1; i < len(exp)-2; i++ {
		if !unicode.IsDigit(exp[i]) {
			if exp[i+1] == '0' && unicode.IsDigit(exp[i+2]) {
				return errors.New("digit should not start with zero")
			}
		}
	}

	return nil
}

func wholeExpressionIsDigit(exp []rune) error {
	for _, char := range exp[:len(exp)-1] {
		if !unicode.IsDigit(char) {
			return nil
		}
	}

	return errors.New("expression should contain operations")
}

func divideByZero(exp []rune) error {
	for i := 0; i < len(exp)-1; i++ {
		if !unicode.IsDigit(exp[i]) {
			if exp[i] == '/' && exp[i+1] == '0' {
				return errors.New("it is forbidden to divide by zero")
			}
		}
	}

	return nil
}

func bracketsAreCorrect(exp []rune) error {
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
