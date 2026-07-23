package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketKsAggregation struct {
	v *types.BucketKsAggregation
}

func NewBucketKsAggregation() *_bucketKsAggregation { _ = "STUB: not implemented"; return nil }

func (s *_bucketKsAggregation) Alternative(alternatives ...string) *_bucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketKsAggregation) Fractions(fractions ...types.Float64) *_bucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketKsAggregation) SamplingMethod(samplingmethod string) *_bucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketKsAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketKsAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketKsAggregation) BucketKsAggregationCaster() *types.BucketKsAggregation {
	_ = "STUB: not implemented"
	return nil
}
