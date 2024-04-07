package app

import (
	"expression_solver/app/agent_components"
	parser2 "expression_solver/app/parser"
	"internal/model"
	"internal/model/expression"
	"log"
	"strconv"
	"time"
)

type ExpressionCreator interface {
	ReadOperation(operationType model.OperationType) (*model.OperationWithDuration, error)
	UpdateExpression(e *expression.Expression) error
	ReadAllExpressionsWithStatus(status expression.Status) ([]*expression.Expression, error)
	CreateExpression(expr *expression.Expression) (int, error)
}

type ExpressionManager struct {
	agent        *agent_components.Agent
	expressionDB ExpressionCreator
}

func NewExpressionManager(agent *agent_components.Agent, expressionDB ExpressionCreator) (*ExpressionManager, error) {
	expressionManager := &ExpressionManager{
		agent:        agent,
		expressionDB: expressionDB,
	}

	expressionsWithStatus, err := expressionDB.ReadAllExpressionsWithStatus(expression.InProcess)
	if err != nil {
		return nil, err
	}

	expressionManager.Init(expressionsWithStatus)

	return expressionManager, nil
}

func (em *ExpressionManager) SaveExpressionAndStartSolve(expressionString string, userId int64) (int, error) {
	exp := expression.NewExpressionInProcess(expressionString, userId)
	id, err := em.expressionDB.CreateExpression(exp)

	if err != nil {
		return 0, err
	}

	go em.ParseExpressionAndSolve(exp)
	return id, nil
}

func (em *ExpressionManager) ParseExpressionAndSolve(exp *expression.Expression) {
	expInfix, err := parser2.TokenizeExpression(exp.Expression)

	if err != nil {
		setInvalidStatus(exp)
		_ = em.expressionDB.UpdateExpression(exp)
		return
	}

	expPostfix, err := parser2.InfixToPostfix(expInfix)

	if err != nil {
		setInvalidStatus(exp)
		_ = em.expressionDB.UpdateExpression(exp)
		return
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

		operationWithDuration, err := em.expressionDB.ReadOperation(token.Operation)
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
			err := em.expressionDB.UpdateExpression(exp)
			if err != nil {
				log.Println(err)
			}
			return
		}

		// Push the result back onto the stack
		stack = append(stack, leastExp.Result)
	}

	if len(stack) < 1 {
		setInvalidStatus(exp)
	} else {
		setResultsToExpression(exp, stack[0])
	}

	err := em.expressionDB.UpdateExpression(exp)
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

func (em *ExpressionManager) Init(expressions []*expression.Expression) {
	for _, exp := range expressions {
		go em.ParseExpressionAndSolve(exp)
	}
}
