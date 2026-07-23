package types

type BucketCorrelationAggregation struct {
	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Function BucketCorrelationFunction `json:"function"`
}

func (s *BucketCorrelationAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketCorrelationAggregation() *BucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}

type BucketCorrelationAggregationVariant interface {
	BucketCorrelationAggregationCaster() *BucketCorrelationAggregation
}

func (s *BucketCorrelationAggregation) BucketCorrelationAggregationCaster() *BucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}
