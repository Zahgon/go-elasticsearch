package types

type TopMetricsAggregation struct {
	Field *string `json:"field,omitempty"`

	Metrics []TopMetricsValue `json:"metrics,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	Size *int `json:"size,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`
}

func (s *TopMetricsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTopMetricsAggregation() *TopMetricsAggregation { _ = "STUB: not implemented"; return nil }

type TopMetricsAggregationVariant interface {
	TopMetricsAggregationCaster() *TopMetricsAggregation
}

func (s *TopMetricsAggregation) TopMetricsAggregationCaster() *TopMetricsAggregation {
	_ = "STUB: not implemented"
	return nil
}
