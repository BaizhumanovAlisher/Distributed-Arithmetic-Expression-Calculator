package grpc_client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"internal/model"
	"internal/model/expression"
	"internal/protos/gen/go/expression_solver_v1"
)

type ExpressionSolver struct {
	client expression_solver_v1.ExpressionSolverClient
}

func NewExpressionSolver(path string) (*ExpressionSolver, error) {
	conn, err := grpc.Dial(path, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := expression_solver_v1.NewExpressionSolverClient(conn)
	return &ExpressionSolver{client: client}, nil
}
func (e *ExpressionSolver) CreateExpressionAndStartSolve(ctx context.Context, expression string, userId int64) (int, error) {
	result, err := e.client.SolveExpression(ctx, &expression_solver_v1.ExpressionRequest{
		Expression: expression,
		UserId:     userId,
	})

	if err != nil {
		return 0, err
	}

	return int(result.GetExpressionId()), nil
}

func (e *ExpressionSolver) GetCalculatorsStatus(ctx context.Context) ([]*expression.MiniCalculator, error) {
	calculatorsFromGrpc, err := e.client.GetCalculatorsStatus(ctx, &expression_solver_v1.Empty{})

	if err != nil {
		return nil, err
	}

	calculators := make([]*expression.MiniCalculator, len(calculatorsFromGrpc.GetCalculators()))
	for i, c := range calculatorsFromGrpc.GetCalculators() {
		calculators[i] = ConvertCalculator(c)
	}

	return calculators, nil
}

func ConvertCalculator(calculator *expression_solver_v1.Calculator) *expression.MiniCalculator {
	return &expression.MiniCalculator{
		ResourceId:      int(calculator.ResourceId),
		LeastExpression: ConvertLeastExpression(calculator.LeastExpression),
	}
}

func ConvertLeastExpression(leastExpression *expression_solver_v1.LeastExpression) *expression.LeastExpression {
	if leastExpression == nil {
		return nil
	}

	operation, _ := model.DefineOperation([]rune(leastExpression.Operator)[0])
	return &expression.LeastExpression{
		Number1:          leastExpression.Number1,
		Number2:          leastExpression.Number2,
		Operation:        operation,
		IdExpression:     int(leastExpression.IdExpression),
		DurationInSecond: int(leastExpression.DurationInSecond),
	}
}
