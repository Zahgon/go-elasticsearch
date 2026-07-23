package types

type BucketMetricValueAggregate struct {
	Keys []string `json:"keys"`
	Meta Metadata `json:"meta,omitempty"`

	Value         *Float64 `json:"value,omitempty"`
	ValueAsString *string  `json:"value_as_string,omitempty"`
}

func (s *BucketMetricValueAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketMetricValueAggregate() *BucketMetricValueAggregate {
	_ = "STUB: not implemented"
	return nil
}
