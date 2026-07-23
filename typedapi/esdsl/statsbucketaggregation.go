package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _statsBucketAggregation struct {
	v *types.StatsBucketAggregation
}

func NewStatsBucketAggregation() *_statsBucketAggregation { _ = "STUB: not implemented"; return nil }

func (s *_statsBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_statsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsBucketAggregation) Format(format string) *_statsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_statsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_statsBucketAggregation) StatsBucketAggregationCaster() *types.StatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
