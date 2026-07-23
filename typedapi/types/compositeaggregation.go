package types

type CompositeAggregation struct {
	After CompositeAggregateKey `json:"after,omitempty"`

	Size *int `json:"size,omitempty"`

	Sources []map[string]CompositeAggregationSource `json:"sources,omitempty"`
}

func (s *CompositeAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeAggregation() *CompositeAggregation { _ = "STUB: not implemented"; return nil }

type CompositeAggregationVariant interface {
	CompositeAggregationCaster() *CompositeAggregation
}

func (s *CompositeAggregation) CompositeAggregationCaster() *CompositeAggregation {
	_ = "STUB: not implemented"
	return nil
}
