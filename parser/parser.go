package parser

import (
	"fmt"
	"unicode"
)

func TokenizeExpression(expression string) ([]*Token, error) {
	var tokens []*Token
	currentNumber := ""
	isNegative := false

	for _, char := range expression {
		if char == '-' && currentNumber == "" {
			isNegative = true
			continue
		}

		token := DefineOneLength(char)

		if token != nil {
			// If there's a current number being built up, create a token for it
			if currentNumber != "" {
				numberStr := currentNumber
				if isNegative {
					numberStr = "-" + numberStr
				}
				tokenPrevious, err := NewTokenFromNumber(numberStr)
				if err != nil {
					return nil, err
				}

				tokens = append(tokens, tokenPrevious)
				currentNumber = ""
				isNegative = false
			}

			tokens = append(tokens, token)
		} else if unicode.IsDigit(char) || char == '.' {
			currentNumber += string(char)
		} else {
			// Unrecognized character
			return nil, fmt.Errorf("unrecognized character: %c", char)
		}
	}

	// Handle any remaining number at the end of the expression
	if currentNumber != "" {
		numberStr := currentNumber
		if isNegative {
			numberStr = "-" + numberStr
		}
		tokenLast, err := NewTokenFromNumber(numberStr)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, tokenLast)
	}

	return tokens, nil
}
