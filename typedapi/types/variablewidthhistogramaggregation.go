package types

type VariableWidthHistogramAggregation struct {
	Buckets *int `json:"buckets,omitempty"`

	Field *string `json:"field,omitempty"`

	InitialBuffer *int    `json:"initial_buffer,omitempty"`
	Script        *Script `json:"script,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`
}

func (s *VariableWidthHistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewVariableWidthHistogramAggregation() *VariableWidthHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

type VariableWidthHistogramAggregationVariant interface {
	VariableWidthHistogramAggregationCaster() *VariableWidthHistogramAggregation
}

func (s *VariableWidthHistogramAggregation) VariableWidthHistogramAggregationCaster() *VariableWidthHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
