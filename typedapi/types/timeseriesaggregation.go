package types

type TimeSeriesAggregation struct {
	Keyed *bool `json:"keyed,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *TimeSeriesAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTimeSeriesAggregation() *TimeSeriesAggregation { _ = "STUB: not implemented"; return nil }

type TimeSeriesAggregationVariant interface {
	TimeSeriesAggregationCaster() *TimeSeriesAggregation
}

func (s *TimeSeriesAggregation) TimeSeriesAggregationCaster() *TimeSeriesAggregation {
	_ = "STUB: not implemented"
	return nil
}
