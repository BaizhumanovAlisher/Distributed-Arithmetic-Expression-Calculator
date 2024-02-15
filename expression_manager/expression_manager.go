package expression_manager

import (
	"distributed_calculator/agent"
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"distributed_calculator/parser"
	"errors"
	"strconv"
	"time"
)

type ReadOperationWithDuration func(operationType model.OperationType) (*model.OperationWithDuration, error)
type UpdateExpression func(*expression.Expression) error
type ReadAllExpressionsWithStatus func(expression.Status) ([]*expression.Expression, error)

type ExpressionManager struct {
	agent *agent.Agent
	ReadOperationWithDuration
	UpdateExpression
	ReadAllExpressionsWithStatus
}

func NewExpressionManager(
	agent *agent.Agent,
	readOperationWithDuration ReadOperationWithDuration,
	updateExpression UpdateExpression,
	readAllExpressionsWithStatus ReadAllExpressionsWithStatus) (*ExpressionManager, error) {

	expressionManager := &ExpressionManager{
		agent:                        agent,
		ReadOperationWithDuration:    readOperationWithDuration,
		UpdateExpression:             updateExpression,
		ReadAllExpressionsWithStatus: readAllExpressionsWithStatus,
	}

	err := expressionManager.Init()

	if err != nil {
		return nil, err
	}

	return expressionManager, err
}

func (em *ExpressionManager) ParseExpressionAndSolve(exp *expression.Expression) error {
	if exp == nil {
		return errors.New("no expression")
	}

	expInfix, err := parser.TokenizeExpression(exp.Expression)

	if err != nil {
		return err
	}

	expPostfix, err := parser.InfixToPostfix(expInfix)

	if err != nil {
		return err
	}

	go em.SolveExpression(exp, expPostfix)

	return nil
}

func (em *ExpressionManager) SolveExpression(exp *expression.Expression, expPostfix []*parser.Token) {
	stack := make([]float64, 0)

	for _, token := range expPostfix {
		if token.Number != nil {
			// Push number onto the stack
			stack = append(stack, *token.Number)
		}

		// Pop two numbers from the stack
		num2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		num1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		operationWihtDuration, err := em.ReadOperationWithDuration(token.Operation)
		var durationInSec int

		if err == nil {
			durationInSec = operationWihtDuration.DurationInSecond
		}

		leastExp := expression.NewLeastExpression(num1, num2, token.Operation, exp.Id, durationInSec)
		em.agent.AddTask(leastExp)

		//Wait until operation will be completed
		resultIsOk := <-leastExp.ResultIsCorrect
		close(leastExp.ResultIsCorrect)

		if !resultIsOk {
			setInvalidStatus(exp)
			em.UpdateExpression(exp)
			return
		}

		// Push the result back onto the stack
		stack = append(stack, leastExp.Result)
	}

	setResultsToExpression(exp, stack[0])
	em.UpdateExpression(exp)
}

func setInvalidStatus(exp *expression.Expression) {
	exp.Status = expression.Invalid
}

func setResultsToExpression(exp *expression.Expression, result float64) {
	exp.Answer = strconv.FormatFloat(result, 'g', 5, 64)
	exp.Status = expression.Completed
	timeCompleted := time.Now()
	exp.CompletedAt = &timeCompleted
}

func (em *ExpressionManager) Init() error {
	expressions, err := em.ReadAllExpressionsWithStatus(expression.InProcess)
	if err != nil {
		return err
	}

	for _, exp := range expressions {
		if err := em.ParseExpressionAndSolve(exp); err != nil {
			setInvalidStatus(exp)
			_ = em.UpdateExpression(exp)
		}
	}

	return nil
}
