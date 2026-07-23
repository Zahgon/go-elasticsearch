package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _movingPercentilesAggregation struct {
	v *types.MovingPercentilesAggregation
}

func NewMovingPercentilesAggregation() *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) Keyed(keyed bool) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) Shift(shift int) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) Window(window int) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) Format(format string) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_movingPercentilesAggregation) MovingPercentilesAggregationCaster() *types.MovingPercentilesAggregation {
	_ = "STUB: not implemented"
	return nil
}
