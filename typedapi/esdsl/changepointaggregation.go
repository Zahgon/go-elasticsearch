package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _changePointAggregation struct {
	v *types.ChangePointAggregation
}

func NewChangePointAggregation() *_changePointAggregation { _ = "STUB: not implemented"; return nil }

func (s *_changePointAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_changePointAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_changePointAggregation) Format(format string) *_changePointAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_changePointAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_changePointAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_changePointAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_changePointAggregation) ChangePointAggregationCaster() *types.ChangePointAggregation {
	_ = "STUB: not implemented"
	return nil
}
