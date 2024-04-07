package parser

import (
	"errors"
	"internal/model"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		name     string
		infix    []*Token
		expected []*Token
		err      error
	}{
		{
			name:     "EmptyExpression",
			infix:    []*Token{},
			expected: []*Token{},
		},
		{
			name:     "NumbersOnly",
			infix:    []*Token{{Number: float64Ptr(1)}, {Number: float64Ptr(2)}, {Number: float64Ptr(3)}},
			expected: []*Token{{Number: float64Ptr(1)}, {Number: float64Ptr(2)}, {Number: float64Ptr(3)}},
		},
		{
			name:     "BasicArithmetic",
			infix:    []*Token{{Number: float64Ptr(1)}, {Operation: model.Addition}, {Number: float64Ptr(2)}, {Operation: model.Multiplication}, {Number: float64Ptr(3)}},
			expected: []*Token{{Number: float64Ptr(1)}, {Number: float64Ptr(2)}, {Number: float64Ptr(3)}, {Operation: model.Multiplication}, {Operation: model.Addition}},
		},
		{
			name:     "Parentheses",
			infix:    []*Token{{Number: float64Ptr(1)}, {Operation: model.Addition}, {IsLeftBracket: true}, {Number: float64Ptr(2)}, {Operation: model.Multiplication}, {Number: float64Ptr(3)}, {IsRightBracket: true}},
			expected: []*Token{{Number: float64Ptr(1)}, {Number: float64Ptr(2)}, {Number: float64Ptr(3)}, {Operation: model.Multiplication}, {Operation: model.Addition}},
		},
		{
			name:     "MixedPrecedence",
			infix:    []*Token{{Number: float64Ptr(1)}, {Operation: model.Addition}, {Number: float64Ptr(2)}, {Operation: model.Multiplication}, {Number: float64Ptr(3)}, {Operation: model.Subtraction}, {Number: float64Ptr(4)}},
			expected: []*Token{{Number: float64Ptr(1)}, {Number: float64Ptr(2)}, {Number: float64Ptr(3)}, {Operation: model.Multiplication}, {Operation: model.Addition}, {Number: float64Ptr(4)}, {Operation: model.Subtraction}},
		},
		{
			name:  "UnmatchedParentheses",
			infix: []*Token{{Number: float64Ptr(1)}, {IsLeftBracket: true}, {Number: float64Ptr(2)}, {Operation: model.Addition}, {Number: float64Ptr(3)}},
			err:   errors.New("mismatched parentheses"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := InfixToPostfix(tt.infix)
			if err != nil {
				if tt.err == nil || err.Error() != tt.err.Error() {
					t.Errorf("InfixToPostfix(%v) unexpected error: got %v, expected %v", tt.infix, err, tt.err)
				}
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("InfixToPostfix(%v) length mismatch: got %v, expected %v", tt.infix, result, tt.expected)
				return
			}

			for i, token := range result {
				if token.Number != nil && tt.expected[i].Number != nil {
					if *token.Number != *tt.expected[i].Number {
						t.Errorf("InfixToPostfix(%v) token mismatch at index %d: got %v, expected %v", tt.infix, i, *token.Number, *tt.expected[i].Number)
					}
				} else if token.Operation != tt.expected[i].Operation {
					t.Errorf("InfixToPostfix(%v) token mismatch at index %d: got %v, expected %v", tt.infix, i, token.Operation, tt.expected[i].Operation)
				} else if token.IsLeftBracket != tt.expected[i].IsLeftBracket {
					t.Errorf("InfixToPostfix(%v) token mismatch at index %d: got %v, expected %v", tt.infix, i, token.IsLeftBracket, tt.expected[i].IsLeftBracket)
				} else if token.IsRightBracket != tt.expected[i].IsRightBracket {
					t.Errorf("InfixToPostfix(%v) token mismatch at index %d: got %v, expected %v", tt.infix, i, token.IsRightBracket, tt.expected[i].IsRightBracket)
				}
			}
		})
	}
}
