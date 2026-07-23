package types

type NestedAggregation struct {
	Path *string `json:"path,omitempty"`
}

func (s *NestedAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNestedAggregation() *NestedAggregation { _ = "STUB: not implemented"; return nil }

type NestedAggregationVariant interface {
	NestedAggregationCaster() *NestedAggregation
}

func (s *NestedAggregation) NestedAggregationCaster() *NestedAggregation {
	_ = "STUB: not implemented"
	return nil
}
