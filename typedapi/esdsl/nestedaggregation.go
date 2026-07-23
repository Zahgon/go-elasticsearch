package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nestedAggregation struct {
	v *types.NestedAggregation
}

func NewNestedAggregation() *_nestedAggregation { _ = "STUB: not implemented"; return nil }

func (s *_nestedAggregation) Path(field string) *_nestedAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nestedAggregation) NestedAggregationCaster() *types.NestedAggregation {
	_ = "STUB: not implemented"
	return nil
}
