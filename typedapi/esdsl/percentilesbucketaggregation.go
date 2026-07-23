package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _percentilesBucketAggregation struct {
	v *types.PercentilesBucketAggregation
}

func NewPercentilesBucketAggregation() *_percentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) Percents(percents ...types.Float64) *_percentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_percentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) Format(format string) *_percentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_percentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_percentilesBucketAggregation) PercentilesBucketAggregationCaster() *types.PercentilesBucketAggregation {
	_ = "STUB: not implemented"
	return nil
}
