package types

type VariableWidthHistogramAggregate struct {
	Buckets BucketsVariableWidthHistogramBucket `json:"buckets"`
	Meta    Metadata                            `json:"meta,omitempty"`
}

func (s *VariableWidthHistogramAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewVariableWidthHistogramAggregate() *VariableWidthHistogramAggregate {
	_ = "STUB: not implemented"
	return nil
}
