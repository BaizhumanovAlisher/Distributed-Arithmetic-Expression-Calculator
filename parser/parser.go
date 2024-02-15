package parser

import (
	"fmt"
	"log/slog"
	"unicode"
)

func TokenizeExpression(expression string) ([]*Token, error) {
	var tokens []*Token
	currentNumber := ""
	isNegative := false

	logger := slog.Default() // Obtain the default logger

	for _, char := range expression {
		if char == '-' && currentNumber == "" {
			isNegative = true
			continue
		}

		token := DefineOneLength(char)

		if token != nil {
			// Log the token creation if needed
			logger.Info("Created token", "value", token.String())

			// If there's a current number being built up, create a token for it
			if currentNumber != "" {
				numberStr := currentNumber
				if isNegative {
					numberStr = "-" + numberStr
				}
				tokenPrevious, err := NewTokenFromNumber(numberStr)
				if err != nil {
					logger.Error("Failed to create number token", "error", err)
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
			err := fmt.Errorf("unrecognized character: %c", char)
			logger.Error("Encountered unrecognized character", "character", char, "error", err)
			return nil, err
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
			logger.Error("Failed to create last number token", "error", err)
			return nil, err
		}
		tokens = append(tokens, tokenLast)
	}

	return tokens, nil
}
