package types

type MissingAggregation struct {
	Field   *string `json:"field,omitempty"`
	Missing Missing `json:"missing,omitempty"`
}

func (s *MissingAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMissingAggregation() *MissingAggregation { _ = "STUB: not implemented"; return nil }

type MissingAggregationVariant interface {
	MissingAggregationCaster() *MissingAggregation
}

func (s *MissingAggregation) MissingAggregationCaster() *MissingAggregation {
	_ = "STUB: not implemented"
	return nil
}
