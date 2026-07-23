package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _extendedStatsBucketAggregation struct {
	v *types.ExtendedStatsBucketAggregation
}

func NewExtendedStatsBucketAggregation() *_extendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) Sigma(sigma types.Float64) *_extendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_extendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) Format(format string) *_extendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_extendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_extendedStatsBucketAggregation) ExtendedStatsBucketAggregationCaster() *types.ExtendedStatsBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
