package parser

import (
	"api_server/internal/model"
	"errors"
)

func InfixToPostfix(infix []*Token) ([]*Token, error) {
	var output []*Token
	var stack []*Token

	for _, token := range infix {
		switch {
		case token.Number != nil:
			// If the token is a number, add it to the output
			output = append(output, token)
		case token.IsLeftBracket:
			// If the token is a left bracket, push it onto the stack
			stack = append(stack, token)
		case token.IsRightBracket:
			// If the token is a right bracket, pop operators from the stack to the output
			// until a left bracket is at the top of the stack
			for len(stack) > 0 && !stack[len(stack)-1].IsLeftBracket {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				return nil, errors.New("mismatched parentheses")
			}
			// Pop the left bracket from the stack
			stack = stack[:len(stack)-1]
		default:
			// If the token is an operator, pop operators from the stack to the output
			// until the top of the stack has an operator of lower precedence or a left bracket
			for len(stack) > 0 && model.Precedence(token.Operation) <= model.Precedence(stack[len(stack)-1].Operation) {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// Push the operator onto the stack
			stack = append(stack, token)
		}
	}

	// Pop any remaining operators from the stack to the output
	for len(stack) > 0 {
		if stack[len(stack)-1].IsLeftBracket {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return output, nil
}
