package types

type GeoCentroidAggregation struct {
	Count *int64 `json:"count,omitempty"`

	Field    *string     `json:"field,omitempty"`
	Location GeoLocation `json:"location,omitempty"`

	Missing Missing `json:"missing,omitempty"`
	Script  *Script `json:"script,omitempty"`
}

func (s *GeoCentroidAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGeoCentroidAggregation() *GeoCentroidAggregation { _ = "STUB: not implemented"; return nil }

type GeoCentroidAggregationVariant interface {
	GeoCentroidAggregationCaster() *GeoCentroidAggregation
}

func (s *GeoCentroidAggregation) GeoCentroidAggregationCaster() *GeoCentroidAggregation {
	_ = "STUB: not implemented"
	return nil
}
