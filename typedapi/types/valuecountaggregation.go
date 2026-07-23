package types

type ValueCountAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *ValueCountAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewValueCountAggregation() *ValueCountAggregation { _ = "STUB: not implemented"; return nil }

type ValueCountAggregationVariant interface {
	ValueCountAggregationCaster() *ValueCountAggregation
}

func (s *ValueCountAggregation) ValueCountAggregationCaster() *ValueCountAggregation {
	_ = "STUB: not implemented"
	return nil
}
