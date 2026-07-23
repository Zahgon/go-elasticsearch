package types

type GeoBoundsAggregation struct {
	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`

	WrapLongitude *bool `json:"wrap_longitude,omitempty"`
}

func (s *GeoBoundsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoBoundsAggregation() *GeoBoundsAggregation { _ = "STUB: not implemented"; return nil }

type GeoBoundsAggregationVariant interface {
	GeoBoundsAggregationCaster() *GeoBoundsAggregation
}

func (s *GeoBoundsAggregation) GeoBoundsAggregationCaster() *GeoBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}
