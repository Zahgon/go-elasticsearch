package types

type GeohexGridAggregation struct {
	Bounds GeoBounds `json:"bounds,omitempty"`

	Field string `json:"field"`

	Precision *int `json:"precision,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *GeohexGridAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeohexGridAggregation() *GeohexGridAggregation { _ = "STUB: not implemented"; return nil }

type GeohexGridAggregationVariant interface {
	GeohexGridAggregationCaster() *GeohexGridAggregation
}

func (s *GeohexGridAggregation) GeohexGridAggregationCaster() *GeohexGridAggregation {
	_ = "STUB: not implemented"
	return nil
}
