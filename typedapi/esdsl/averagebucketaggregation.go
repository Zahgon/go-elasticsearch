package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _averageBucketAggregation struct {
	v *types.AverageBucketAggregation
}

func NewAverageBucketAggregation() *_averageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_averageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageBucketAggregation) Format(format string) *_averageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_averageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_averageBucketAggregation) AverageBucketAggregationCaster() *types.AverageBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
