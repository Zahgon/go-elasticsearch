package types

type AggregationRange struct {
	From *Float64 `json:"from,omitempty"`

	Key *string `json:"key,omitempty"`

	To *Float64 `json:"to,omitempty"`
}

func (s *AggregationRange) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAggregationRange() *AggregationRange { _ = "STUB: not implemented"; return nil }

type AggregationRangeVariant interface {
	AggregationRangeCaster() *AggregationRange
}

func (s *AggregationRange) AggregationRangeCaster() *AggregationRange {
	_ = "STUB: not implemented"
	return nil
}
