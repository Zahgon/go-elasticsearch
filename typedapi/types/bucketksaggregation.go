package types

type BucketKsAggregation struct {
	Alternative []string `json:"alternative,omitempty"`

	BucketsPath BucketsPath `json:"buckets_path,omitempty"`

	Fractions []Float64 `json:"fractions,omitempty"`

	SamplingMethod *string `json:"sampling_method,omitempty"`
}

func (s *BucketKsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBucketKsAggregation() *BucketKsAggregation { _ = "STUB: not implemented"; return nil }

type BucketKsAggregationVariant interface {
	BucketKsAggregationCaster() *BucketKsAggregation
}

func (s *BucketKsAggregation) BucketKsAggregationCaster() *BucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}
