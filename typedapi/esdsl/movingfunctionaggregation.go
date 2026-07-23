package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _movingFunctionAggregation struct {
	v *types.MovingFunctionAggregation
}

func NewMovingFunctionAggregation() *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) Script(script string) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) Shift(shift int) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) Window(window int) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) Format(format string) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingFunctionAggregation) MovingFunctionAggregationCaster() *types.MovingFunctionAggregation {
	_ = "STUB: not implemented"
	return nil
}
