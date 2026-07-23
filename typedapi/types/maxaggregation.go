package types

type MaxAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *MaxAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMaxAggregation() *MaxAggregation { _ = "STUB: not implemented"; return nil }

type MaxAggregationVariant interface {
	MaxAggregationCaster() *MaxAggregation
}

func (s *MaxAggregation) MaxAggregationCaster() *MaxAggregation {
	_ = "STUB: not implemented"
	return nil
}
