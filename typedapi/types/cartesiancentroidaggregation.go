package types

type CartesianCentroidAggregation struct {
	Field *string `json:"field,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *CartesianCentroidAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCartesianCentroidAggregation() *CartesianCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CartesianCentroidAggregationVariant interface {
	CartesianCentroidAggregationCaster() *CartesianCentroidAggregation
}

func (s *CartesianCentroidAggregation) CartesianCentroidAggregationCaster() *CartesianCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}
