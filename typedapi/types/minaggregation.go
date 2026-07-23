package types

type MinAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *MinAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMinAggregation() *MinAggregation { _ = "STUB: not implemented"; return nil }

type MinAggregationVariant interface {
	MinAggregationCaster() *MinAggregation
}

func (s *MinAggregation) MinAggregationCaster() *MinAggregation {
	_ = "STUB: not implemented"
	return nil
}
