package types

type FrequentItemSetsAggregation struct {
	Fields []FrequentItemSetsField `json:"fields"`

	Filter *Query `json:"filter,omitempty"`

	MinimumSetSize *int `json:"minimum_set_size,omitempty"`

	MinimumSupport *Float64 `json:"minimum_support,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *FrequentItemSetsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFrequentItemSetsAggregation() *FrequentItemSetsAggregation {
	_ = "STUB: not implemented"
	return nil
}

type FrequentItemSetsAggregationVariant interface {
	FrequentItemSetsAggregationCaster() *FrequentItemSetsAggregation
}

func (s *FrequentItemSetsAggregation) FrequentItemSetsAggregationCaster() *FrequentItemSetsAggregation {
	_ = "STUB: not implemented"
	return nil
}
