package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _simpleMovingAverageAggregation struct {
	v *types.SimpleMovingAverageAggregation
}

func NewSimpleMovingAverageAggregation(settings types.EmptyObjectVariant) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) Settings(settings types.EmptyObjectVariant) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) Format(format string) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) Minimize(minimize bool) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) Predict(predict int) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) Window(window int) *_simpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_simpleMovingAverageAggregation) SimpleMovingAverageAggregationCaster() *types.SimpleMovingAverageAggregation {
	_ = "STUB: not implemented"
	return nil
}
