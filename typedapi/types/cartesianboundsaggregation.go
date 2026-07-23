package types

type CartesianBoundsAggregation struct {
	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *CartesianBoundsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCartesianBoundsAggregation() *CartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CartesianBoundsAggregationVariant interface {
	CartesianBoundsAggregationCaster() *CartesianBoundsAggregation
}

func (s *CartesianBoundsAggregation) CartesianBoundsAggregationCaster() *CartesianBoundsAggregation {
	_ = "STUB: not implemented"
	return nil
}
