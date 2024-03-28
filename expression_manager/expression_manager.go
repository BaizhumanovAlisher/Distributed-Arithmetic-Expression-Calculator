package expression_manager

import (
	"distributed_calculator/expression_manager/agent"
	parser2 "distributed_calculator/expression_manager/parser"
	"distributed_calculator/model"
	"distributed_calculator/model/expression"
	"log"
	"strconv"
	"time"
)

type ReadOperationWithDuration func(operationType model.OperationType) (*model.OperationWithDuration, error)
type UpdateExpression func(*expression.Expression) error

type ExpressionManager struct {
	agent *agent.Agent
	ReadOperationWithDuration
	UpdateExpression
}

func NewExpressionManager(
	agent *agent.Agent,
	readOperationWithDuration ReadOperationWithDuration,
	updateExpression UpdateExpression,
	readAllExpressionsWithStatus func(expression.Status) ([]*expression.Expression, error)) (*ExpressionManager, error) {

	expressionManager := &ExpressionManager{
		agent:                     agent,
		ReadOperationWithDuration: readOperationWithDuration,
		UpdateExpression:          updateExpression,
	}

	err := expressionManager.Init(readAllExpressionsWithStatus)

	if err != nil {
		return nil, err
	}

	return expressionManager, nil
}

func (em *ExpressionManager) ParseExpressionAndSolve(exp *expression.Expression) {
	if exp == nil {
		log.Println("exp is nil")
	}

	expInfix, err := parser2.TokenizeExpression(exp.Expression)

	if err != nil {
		setInvalidStatus(exp)
		em.UpdateExpression(exp)
	}

	expPostfix, err := parser2.InfixToPostfix(expInfix)

	if err != nil {
		setInvalidStatus(exp)
		em.UpdateExpression(exp)
	}

	em.SolveExpression(exp, expPostfix)
}

func (em *ExpressionManager) SolveExpression(exp *expression.Expression, expPostfix []*parser2.Token) {
	stack := make([]float64, 0)

	for _, token := range expPostfix {
		if token.Number != nil {
			// Push number onto the stack
			stack = append(stack, *token.Number)
			continue
		}

		// Pop two numbers from the stack
		if len(stack) < 2 {
			continue
		}
		num2 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		num1 := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		operationWithDuration, err := em.ReadOperationWithDuration(token.Operation)
		var durationInSec int

		if err == nil {
			durationInSec = operationWithDuration.DurationInSecond
		}

		leastExp := expression.NewLeastExpression(num1, num2, token.Operation, exp.Id, durationInSec)
		em.agent.AddTask(leastExp)

		//Wait until operation will be completed
		resultIsOk := <-leastExp.ResultIsCorrect
		close(leastExp.ResultIsCorrect)

		if !resultIsOk {
			setInvalidStatus(exp)
			err := em.UpdateExpression(exp)
			if err != nil {
				log.Println(err)
			}
			return
		}

		// Push the result back onto the stack
		stack = append(stack, leastExp.Result)
	}

	setResultsToExpression(exp, stack[0])
	err := em.UpdateExpression(exp)
	if err != nil {
		log.Println(err)
	}
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

func (em *ExpressionManager) Init(readAllExpressionsWithStatus func(expression.Status) ([]*expression.Expression, error)) error {
	expressions, err := readAllExpressionsWithStatus(expression.InProcess)
	if err != nil {
		return err
	}

	for _, exp := range expressions {
		go em.ParseExpressionAndSolve(exp)
	}

	return nil
}

func (em *ExpressionManager) StartSolveConcurrently(exp *expression.Expression) {
	newExpression := *exp
	go em.ParseExpressionAndSolve(&newExpression)
}
