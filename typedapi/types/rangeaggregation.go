package types

type RangeAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	Missing *int `json:"missing,omitempty"`

	Ranges []AggregationRange `json:"ranges,omitempty"`
	Script *Script            `json:"script,omitempty"`
}

func (s *RangeAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRangeAggregation() *RangeAggregation { _ = "STUB: not implemented"; return nil }

type RangeAggregationVariant interface {
	RangeAggregationCaster() *RangeAggregation
}

func (s *RangeAggregation) RangeAggregationCaster() *RangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
