package types

type HistogramAggregation struct {
	ExtendedBounds *ExtendedBoundsdouble `json:"extended_bounds,omitempty"`

	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	HardBounds *ExtendedBoundsdouble `json:"hard_bounds,omitempty"`

	Interval *Float64 `json:"interval,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	MinDocCount *int `json:"min_doc_count,omitempty"`

	Missing *Float64 `json:"missing,omitempty"`

	Offset *Float64 `json:"offset,omitempty"`

	Order  AggregateOrder `json:"order,omitempty"`
	Script *Script        `json:"script,omitempty"`
}

func (s *HistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHistogramAggregation() *HistogramAggregation { _ = "STUB: not implemented"; return nil }

type HistogramAggregationVariant interface {
	HistogramAggregationCaster() *HistogramAggregation
}

func (s *HistogramAggregation) HistogramAggregationCaster() *HistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
