package types

type IpRangeAggregation struct {
	Field *string `json:"field,omitempty"`

	Ranges []IpRangeAggregationRange `json:"ranges,omitempty"`
}

func (s *IpRangeAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIpRangeAggregation() *IpRangeAggregation { _ = "STUB: not implemented"; return nil }

type IpRangeAggregationVariant interface {
	IpRangeAggregationCaster() *IpRangeAggregation
}

func (s *IpRangeAggregation) IpRangeAggregationCaster() *IpRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
