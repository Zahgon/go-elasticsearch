package types

type IpRangeAggregationRange struct {
	From *string `json:"from,omitempty"`

	Mask *string `json:"mask,omitempty"`

	To *string `json:"to,omitempty"`
}

func (s *IpRangeAggregationRange) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIpRangeAggregationRange() *IpRangeAggregationRange { _ = "STUB: not implemented"; return nil }

type IpRangeAggregationRangeVariant interface {
	IpRangeAggregationRangeCaster() *IpRangeAggregationRange
}

func (s *IpRangeAggregationRange) IpRangeAggregationRangeCaster() *IpRangeAggregationRange {
	_ = "STUB: not implemented"
	return nil
}
