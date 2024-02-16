package parser

import (
	"distributed_calculator/model"
	"errors"
	"strconv"
)

var errEmptyString = errors.New("empty string")
var errProblemParsing = errors.New("problem parsing")

type Token struct {
	Number         *float64
	Operation      model.OperationType
	IsLeftBracket  bool
	IsRightBracket bool
}

func (t Token) String() string {
	switch {
	case t.Number != nil:
		return strconv.FormatFloat(*t.Number, 'f', -1, 64)
	case t.IsLeftBracket:
		return "("
	case t.IsRightBracket:
		return ")"
	default:
		return string(t.Operation)
	}
}

func DefineOneLength(s rune) *Token {
	switch s {
	case '(':
		return &Token{IsLeftBracket: true}
	case ')':
		return &Token{IsRightBracket: true}
	default:
		operation, ok := model.DefineOperation(s)

		if !ok {
			return nil
		}

		return &Token{Operation: operation}
	}
}

func NewTokenFromNumber(s string) (*Token, error) {
	if len(s) == 0 {
		return nil, errEmptyString
	}

	number, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, errProblemParsing
	}
	return &Token{Number: &number}, nil
}

func float64Ptr(f float64) *float64 {
	return &f
}
