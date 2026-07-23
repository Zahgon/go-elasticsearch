package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _cumulativeSumAggregation struct {
	v *types.CumulativeSumAggregation
}

func NewCumulativeSumAggregation() *_cumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeSumAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_cumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeSumAggregation) Format(format string) *_cumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeSumAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_cumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeSumAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cumulativeSumAggregation) CumulativeSumAggregationCaster() *types.CumulativeSumAggregation {
	_ = "STUB: not implemented"
	return nil
}
