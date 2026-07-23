package types

type ParentAggregation struct {
	Type *string `json:"type,omitempty"`
}

func (s *ParentAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewParentAggregation() *ParentAggregation { _ = "STUB: not implemented"; return nil }

type ParentAggregationVariant interface {
	ParentAggregationCaster() *ParentAggregation
}

func (s *ParentAggregation) ParentAggregationCaster() *ParentAggregation {
	_ = "STUB: not implemented"
	return nil
}
