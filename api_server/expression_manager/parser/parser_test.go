package parser

import (
	"internal/model"
	"testing"
)

func TestTokenizeExpression(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []*Token
	}{
		{
			name:     "EmptyString",
			input:    "",
			expected: []*Token{},
		},
		{
			name:  "MixedDigitsAndOperators",
			input: "12+34*5",
			expected: []*Token{
				{Number: float64Ptr(12)},
				{Operation: model.Addition},
				{Number: float64Ptr(34)},
				{Operation: model.Multiplication},
				{Number: float64Ptr(5)},
			},
		},
		{
			name:  "Parentheses",
			input: "(1+2)*3",
			expected: []*Token{
				{IsLeftBracket: true},
				{Number: float64Ptr(1)},
				{Operation: model.Addition},
				{Number: float64Ptr(2)},
				{IsRightBracket: true},
				{Operation: model.Multiplication},
				{Number: float64Ptr(3)},
			},
		},
		{
			name:  "_",
			input: "12+34*5",
			expected: []*Token{
				{Number: float64Ptr(12)},
				{Operation: model.Addition},
				{Number: float64Ptr(34)},
				{Operation: model.Multiplication},
				{Number: float64Ptr(5)},
			},
		},
		{
			name:     "InvalidCharacter",
			input:    "12a+34*5",
			expected: nil,
		},
		{
			name:  "NegativeNumber",
			input: "-12+34*5",
			expected: []*Token{
				{Number: float64Ptr(-12)},
				{Operation: model.Addition},
				{Number: float64Ptr(34)},
				{Operation: model.Multiplication},
				{Number: float64Ptr(5)},
			},
		},
		{
			name:  "FloatingPointNumber",
			input: "12.34+5.67*8.90",
			expected: []*Token{
				{Number: float64Ptr(12.34)},
				{Operation: model.Addition},
				{Number: float64Ptr(5.67)},
				{Operation: model.Multiplication},
				{Number: float64Ptr(8.90)},
			},
		},
		{
			name:  "ComplexExpression",
			input: "((1+2)*3)/(4-5)",
			expected: []*Token{
				{IsLeftBracket: true},
				{IsLeftBracket: true},
				{Number: float64Ptr(1)},
				{Operation: model.Addition},
				{Number: float64Ptr(2)},
				{IsRightBracket: true},
				{Operation: model.Multiplication},
				{Number: float64Ptr(3)},
				{IsRightBracket: true},
				{Operation: model.Division},
				{IsLeftBracket: true},
				{Number: float64Ptr(4)},
				{Operation: model.Subtraction},
				{Number: float64Ptr(5)},
				{IsRightBracket: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TokenizeExpression(tt.input)
			if err != nil {
				if tt.expected != nil {
					t.Errorf("TokenizeExpression(%q) unexpected error: %v", tt.input, err)
				}
				return
			}

			if len(result) != len(tt.expected) {
				t.Errorf("TokenizeExpression(%q) length mismatch: got %d, expected %d", tt.input, len(result), len(tt.expected))
				return
			}

			for i, token := range result {
				if token.Number != nil && tt.expected[i].Number != nil {
					if *token.Number != *tt.expected[i].Number {
						t.Errorf("TokenizeExpression(%q) token mismatch at index %d: got %v, expected %v", tt.input, i, *token.Number, *tt.expected[i].Number)
					}
				} else if token.Operation != tt.expected[i].Operation {
					t.Errorf("TokenizeExpression(%q) token mismatch at index %d: got %v, expected %v", tt.input, i, token.Operation, tt.expected[i].Operation)
				} else if token.IsLeftBracket != tt.expected[i].IsLeftBracket {
					t.Errorf("TokenizeExpression(%q) token mismatch at index %d: got %v, expected %v", tt.input, i, token.IsLeftBracket, tt.expected[i].IsLeftBracket)
				} else if token.IsRightBracket != tt.expected[i].IsRightBracket {
					t.Errorf("TokenizeExpression(%q) token mismatch at index %d: got %v, expected %v", tt.input, i, token.IsRightBracket, tt.expected[i].IsRightBracket)
				}
			}
		})
	}
}
