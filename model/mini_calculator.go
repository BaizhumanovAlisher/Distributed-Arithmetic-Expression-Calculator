package model

type MiniCalculator struct {
	ResourceId      int             `json:"resourceId"`
	LeastExpression LeastExpression `json:"leastExpression"`
}

func NewMiniCalculator() *MiniCalculator {
	return &MiniCalculator{}
}
