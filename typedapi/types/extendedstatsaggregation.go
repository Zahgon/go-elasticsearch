package types

type ExtendedStatsAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	Sigma *Float64 `json:"sigma,omitempty"`
}

func (s *ExtendedStatsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExtendedStatsAggregation() *ExtendedStatsAggregation { _ = "STUB: not implemented"; return nil }

type ExtendedStatsAggregationVariant interface {
	ExtendedStatsAggregationCaster() *ExtendedStatsAggregation
}

func (s *ExtendedStatsAggregation) ExtendedStatsAggregationCaster() *ExtendedStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
