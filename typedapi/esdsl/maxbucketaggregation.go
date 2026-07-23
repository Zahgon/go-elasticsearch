package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _maxBucketAggregation struct {
	v *types.MaxBucketAggregation
}

func NewMaxBucketAggregation() *_maxBucketAggregation { _ = "STUB: not implemented"; return nil }

func (s *_maxBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_maxBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxBucketAggregation) Format(format string) *_maxBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_maxBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_maxBucketAggregation) MaxBucketAggregationCaster() *types.MaxBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
