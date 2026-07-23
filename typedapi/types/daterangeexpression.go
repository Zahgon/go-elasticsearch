package types

type DateRangeExpression struct {
	From FieldDateMath `json:"from,omitempty"`

	Key *string `json:"key,omitempty"`

	To FieldDateMath `json:"to,omitempty"`
}

func (s *DateRangeExpression) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeExpression() *DateRangeExpression { _ = "STUB: not implemented"; return nil }

type DateRangeExpressionVariant interface {
	DateRangeExpressionCaster() *DateRangeExpression
}

func (s *DateRangeExpression) DateRangeExpressionCaster() *DateRangeExpression {
	_ = "STUB: not implemented"
	return nil
}
