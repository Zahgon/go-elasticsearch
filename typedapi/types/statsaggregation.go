package types

type StatsAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *StatsAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStatsAggregation() *StatsAggregation { _ = "STUB: not implemented"; return nil }

type StatsAggregationVariant interface {
	StatsAggregationCaster() *StatsAggregation
}

func (s *StatsAggregation) StatsAggregationCaster() *StatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
