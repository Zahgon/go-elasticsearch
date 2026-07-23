package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _sumBucketAggregation struct {
	v *types.SumBucketAggregation
}

func NewSumBucketAggregation() *_sumBucketAggregation { _ = "STUB: not implemented"; return nil }

func (s *_sumBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_sumBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumBucketAggregation) Format(format string) *_sumBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_sumBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sumBucketAggregation) SumBucketAggregationCaster() *types.SumBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
