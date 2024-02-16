package validators

import (
	"errors"
	"testing"
)

func TestStartWithNumber(t *testing.T) {
	testCases := []struct {
		name     string
		exp      []rune
		expected error
	}{
		{
			name:     "Starts with number",
			exp:      []rune("1+2"),
			expected: nil,
		},
		{
			name:     "Does not start with number",
			exp:      []rune("a+2"),
			expected: errors.New("expression should start with number"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := startWithNumber(tc.exp)
			if (got == nil) != (tc.expected == nil) {
				t.Errorf("startWithNumber() = %v, want %v", got, tc.expected)
			}
		})
	}
}

func TestNotContainExtraCharacters(t *testing.T) {
	tests := []struct {
		name string
		exp  []rune
		want error
	}{
		{
			name: "Valid expression",
			exp:  []rune("(1+2)"),
			want: nil,
		},
		{
			name: "Expression with extra characters",
			exp:  []rune("(1+2)!"),
			want: errors.New("expression should not contain extra symbols"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := notContainExtraCharacters(tt.exp)
			if (err != nil) != (tt.want != nil) {
				t.Errorf("notContainExtraCharacters() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestDigitStartWithZero(t *testing.T) {
	tests := []struct {
		name string
		exp  []rune
		want error
	}{
		{
			name: "Valid expression",
			exp:  []rune("1+2"),
			want: nil,
		},
		{
			name: "Expression starting with zero",
			exp:  []rune("01+2"),
			want: errors.New("digit should not start with zero"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := digitStartWithZero(tt.exp)
			if (err != nil) != (tt.want != nil) {
				t.Errorf("digitStartWithZero() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestWholeExpressionIsDigit(t *testing.T) {
	tests := []struct {
		name string
		exp  []rune
		want error
	}{
		{
			name: "All digits",
			exp:  []rune("12345"),
			want: errors.New("expression should contain operations"),
		},
		{
			name: "Mixed digits and operations",
			exp:  []rune("12+34"),
			want: nil,
		},
		{
			name: "Only operations",
			exp:  []rune("+-*/"),
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := wholeExpressionIsDigit(tt.exp)
			if (err != nil) != (tt.want != nil) {
				t.Errorf("wholeExpressionIsDigit() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}

func TestBracketsAreCorrect(t *testing.T) {
	tests := []struct {
		name string
		exp  []rune
		want error
	}{
		{
			name: "Correct brackets",
			exp:  []rune("(1+2)"),
			want: nil,
		},
		{
			name: "Missing opening bracket",
			exp:  []rune("1+2)"),
			want: errors.New("brackets should be correct"),
		},
		{
			name: "Missing closing bracket",
			exp:  []rune("(1+2"),
			want: errors.New("brackets should be correct"),
		},
		{
			name: "Mismatched brackets",
			exp:  []rune("(1+2]"),
			want: errors.New("brackets should be correct"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := bracketsAreCorrect(tt.exp)
			if (err != nil) != (tt.want != nil) {
				t.Errorf("bracketsAreCorrect() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
