package expression

type MiniCalculator struct {
	ResourceId      int              `json:"resourceId"`
	LeastExpression *LeastExpression `json:"leastExpression"`
}

func NewMiniCalculator(resourceId int) *MiniCalculator {
	return &MiniCalculator{ResourceId: resourceId}
}
