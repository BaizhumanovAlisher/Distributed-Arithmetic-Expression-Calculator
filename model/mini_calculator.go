package model

import "distributed_calculator/model/expression"

type MiniCalculator struct {
	ResourceId      int                         `json:"resourceId"`
	LeastExpression *expression.LeastExpression `json:"leastExpression"`
}

func NewMiniCalculator(resourceId int) *MiniCalculator {
	return &MiniCalculator{ResourceId: resourceId}
}
