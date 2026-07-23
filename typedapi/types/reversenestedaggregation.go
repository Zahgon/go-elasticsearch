package types

type ReverseNestedAggregation struct {
	Path *string `json:"path,omitempty"`
}

func (s *ReverseNestedAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReverseNestedAggregation() *ReverseNestedAggregation { _ = "STUB: not implemented"; return nil }

type ReverseNestedAggregationVariant interface {
	ReverseNestedAggregationCaster() *ReverseNestedAggregation
}

func (s *ReverseNestedAggregation) ReverseNestedAggregationCaster() *ReverseNestedAggregation {
	_ = "STUB: not implemented"
	return nil
}
