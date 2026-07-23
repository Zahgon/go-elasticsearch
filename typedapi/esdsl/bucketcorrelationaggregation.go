package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketCorrelationAggregation struct {
	v *types.BucketCorrelationAggregation
}

func NewBucketCorrelationAggregation(function types.BucketCorrelationFunctionVariant) *_bucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationAggregation) Function(function types.BucketCorrelationFunctionVariant) *_bucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bucketCorrelationAggregation) BucketCorrelationAggregationCaster() *types.BucketCorrelationAggregation {
	_ = "STUB: not implemented"
	return nil
}
