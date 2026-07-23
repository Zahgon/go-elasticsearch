package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _globalAggregation struct {
	v *types.GlobalAggregation
}

func NewGlobalAggregation() *_globalAggregation { _ = "STUB: not implemented"; return nil }

func (s *_globalAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_globalAggregation) GlobalAggregationCaster() *types.GlobalAggregation {
	_ = "STUB: not implemented"
	return nil
}
