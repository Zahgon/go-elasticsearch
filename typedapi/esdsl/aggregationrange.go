package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _aggregationRange struct {
	v *types.AggregationRange
}

func NewAggregationRange() *_aggregationRange { _ = "STUB: not implemented"; return nil }

func (s *_aggregationRange) From(from types.Float64) *_aggregationRange {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregationRange) Key(key string) *_aggregationRange {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregationRange) To(to types.Float64) *_aggregationRange {
	_ = "STUB: not implemented"
	return nil
}

func (s *_aggregationRange) AggregationRangeCaster() *types.AggregationRange {
	_ = "STUB: not implemented"
	return nil
}
