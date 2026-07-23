package types

type ChildrenAggregation struct {
	Type *string `json:"type,omitempty"`
}

func (s *ChildrenAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewChildrenAggregation() *ChildrenAggregation { _ = "STUB: not implemented"; return nil }

type ChildrenAggregationVariant interface {
	ChildrenAggregationCaster() *ChildrenAggregation
}

func (s *ChildrenAggregation) ChildrenAggregationCaster() *ChildrenAggregation {
	_ = "STUB: not implemented"
	return nil
}
