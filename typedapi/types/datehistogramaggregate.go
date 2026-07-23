package types

type DateHistogramAggregate struct {
	Buckets BucketsDateHistogramBucket `json:"buckets"`
	Meta    Metadata                   `json:"meta,omitempty"`
}

func (s *DateHistogramAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateHistogramAggregate() *DateHistogramAggregate { _ = "STUB: not implemented"; return nil }
