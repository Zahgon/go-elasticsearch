package types

type StringStatsAggregation struct {
	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	ShowDistribution *bool `json:"show_distribution,omitempty"`
}

func (s *StringStatsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStringStatsAggregation() *StringStatsAggregation { _ = "STUB: not implemented"; return nil }

type StringStatsAggregationVariant interface {
	StringStatsAggregationCaster() *StringStatsAggregation
}

func (s *StringStatsAggregation) StringStatsAggregationCaster() *StringStatsAggregation {
	_ = "STUB: not implemented"
	return nil
}
