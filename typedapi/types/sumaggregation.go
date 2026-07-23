package types

type SumAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *SumAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSumAggregation() *SumAggregation { _ = "STUB: not implemented"; return nil }

type SumAggregationVariant interface {
	SumAggregationCaster() *SumAggregation
}

func (s *SumAggregation) SumAggregationCaster() *SumAggregation {
	_ = "STUB: not implemented"
	return nil
}
