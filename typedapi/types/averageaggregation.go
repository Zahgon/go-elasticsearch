package types

type AverageAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *AverageAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAverageAggregation() *AverageAggregation { _ = "STUB: not implemented"; return nil }

type AverageAggregationVariant interface {
	AverageAggregationCaster() *AverageAggregation
}

func (s *AverageAggregation) AverageAggregationCaster() *AverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
