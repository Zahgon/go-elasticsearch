package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _minBucketAggregation struct {
	v *types.MinBucketAggregation
}

func NewMinBucketAggregation() *_minBucketAggregation { _ = "STUB: not implemented"; return nil }

func (s *_minBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_minBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minBucketAggregation) Format(format string) *_minBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_minBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_minBucketAggregation) MinBucketAggregationCaster() *types.MinBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
