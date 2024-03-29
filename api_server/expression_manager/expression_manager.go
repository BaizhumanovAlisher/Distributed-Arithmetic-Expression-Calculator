package expression_manager

import (
	"api_server/expression_manager/agent"
	parser2 "api_server/expression_manager/parser"
	"api_server/internal/model"
	expression2 "api_server/internal/model/expression"
	"api_server/internal/storage/postgresql"
	"log"
	"strconv"
	"time"
)

type ReadOperationWithDuration func(operationType model.OperationType) (*model.OperationWithDuration, error)
type UpdateExpression func(*expression2.Expression) error

type ExpressionManager struct {
	agent *agent.Agent
	ReadOperationWithDuration
	UpdateExpression
}

func NewExpressionManager(agent *agent.Agent, repo *postgresql.PostgresqlDB) (*ExpressionManager, error) {
	expressionManager := &ExpressionManager{
		agent:                     agent,
		ReadOperationWithDuration: repo.ReadOperation,
		UpdateExpression:          repo.UpdateExpression,
	}

	err := expressionManager.Init(repo.ReadAllExpressionsWithStatus)

	if err != nil {
		return nil, err
	}

	return expressionManager, nil
}

func (em *ExpressionManager) ParseExpressionAndSolve(exp *expression2.Expression) {
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

func (em *ExpressionManager) SolveExpression(exp *expression2.Expression, expPostfix []*parser2.Token) {
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

		leastExp := expression2.NewLeastExpression(num1, num2, token.Operation, exp.Id, durationInSec)
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

func setInvalidStatus(exp *expression2.Expression) {
	exp.Status = expression2.Invalid
}

func setResultsToExpression(exp *expression2.Expression, result float64) {
	exp.Answer = strconv.FormatFloat(result, 'g', 5, 64)
	exp.Status = expression2.Completed
	timeCompleted := time.Now()
	exp.CompletedAt = &timeCompleted
}

func (em *ExpressionManager) Init(readAllExpressionsWithStatus func(expression2.Status) ([]*expression2.Expression, error)) error {
	expressions, err := readAllExpressionsWithStatus(expression2.InProcess)
	if err != nil {
		return err
	}

	for _, exp := range expressions {
		go em.ParseExpressionAndSolve(exp)
	}

	return nil
}

func (em *ExpressionManager) StartSolveConcurrently(exp *expression2.Expression) {
	newExpression := *exp
	go em.ParseExpressionAndSolve(&newExpression)
}
