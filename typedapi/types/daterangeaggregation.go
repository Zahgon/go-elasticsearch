package types

type DateRangeAggregation struct {
	Field *string `json:"field,omitempty"`

	Format *string `json:"format,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Ranges []DateRangeExpression `json:"ranges,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *DateRangeAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateRangeAggregation() *DateRangeAggregation { _ = "STUB: not implemented"; return nil }

type DateRangeAggregationVariant interface {
	DateRangeAggregationCaster() *DateRangeAggregation
}

func (s *DateRangeAggregation) DateRangeAggregationCaster() *DateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
