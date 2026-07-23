package types

type ProjectRoutingExpression struct {
	Expression string `json:"expression"`
}

func (s *ProjectRoutingExpression) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewProjectRoutingExpression() *ProjectRoutingExpression { _ = "STUB: not implemented"; return nil }

type ProjectRoutingExpressionVariant interface {
	ProjectRoutingExpressionCaster() *ProjectRoutingExpression
}

func (s *ProjectRoutingExpression) ProjectRoutingExpressionCaster() *ProjectRoutingExpression {
	_ = "STUB: not implemented"
	return nil
}
