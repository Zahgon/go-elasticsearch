package types

type AutoDateHistogramAggregate struct {
	Buckets  BucketsDateHistogramBucket `json:"buckets"`
	Interval string                     `json:"interval"`
	Meta     Metadata                   `json:"meta,omitempty"`
}

func (s *AutoDateHistogramAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoDateHistogramAggregate() *AutoDateHistogramAggregate {
	_ = "STUB: not implemented"
	return nil
}
