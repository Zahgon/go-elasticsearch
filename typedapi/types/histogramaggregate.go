package types

type HistogramAggregate struct {
	Buckets BucketsHistogramBucket `json:"buckets"`
	Meta    Metadata               `json:"meta,omitempty"`
}

func (s *HistogramAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewHistogramAggregate() *HistogramAggregate { _ = "STUB: not implemented"; return nil }
