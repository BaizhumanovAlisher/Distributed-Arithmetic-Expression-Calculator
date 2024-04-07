package app

import (
	"context"
	"expression_solver/app/agent_components"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"internal/model/expression"
	controller "internal/protos/gen/go/expression_solver_v1"
)

type GrpcController struct {
	controller.UnimplementedExpressionSolverServer
	agent             *agent_components.Agent
	expressionManager *ExpressionManager
}

func NewGrpcController(agent *agent_components.Agent, expressionManager *ExpressionManager) *GrpcController {
	return &GrpcController{
		agent:             agent,
		expressionManager: expressionManager,
	}
}

func (g *GrpcController) SolveExpression(_ context.Context, request *controller.ExpressionRequest) (*controller.ExpressionResponse, error) {
	id, err := g.expressionManager.SaveExpressionAndStartSolve(request.Expression, request.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &controller.ExpressionResponse{ExpressionId: int32(id)}, nil
}

func (g *GrpcController) GetCalculatorsStatus(_ context.Context, _ *controller.Empty) (*controller.CalculatorList, error) {
	miniCalculators := g.agent.GetAllMiniCalculators()
	calculators := make([]*controller.Calculator, len(miniCalculators))

	for i, miniCalculator := range miniCalculators {
		calculators[i] = ConvertCalculator(miniCalculator)
	}

	return &controller.CalculatorList{Calculators: calculators}, nil
}

func ConvertCalculator(calculator *expression.MiniCalculator) *controller.Calculator {
	return &controller.Calculator{
		ResourceId:      int32(calculator.ResourceId),
		LeastExpression: ConvertLeastExpression(calculator.LeastExpression),
	}
}

func ConvertLeastExpression(leastExpression *expression.LeastExpression) *controller.LeastExpression {
	if leastExpression == nil {
		return nil
	}

	return &controller.LeastExpression{
		Number1:          leastExpression.Number1,
		Number2:          leastExpression.Number2,
		Operator:         string(leastExpression.Operation),
		IdExpression:     int32(leastExpression.IdExpression),
		DurationInSecond: int32(leastExpression.DurationInSecond),
	}
}
