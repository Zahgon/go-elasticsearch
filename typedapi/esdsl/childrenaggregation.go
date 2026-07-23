package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _childrenAggregation struct {
	v *types.ChildrenAggregation
}

func NewChildrenAggregation() *_childrenAggregation { _ = "STUB: not implemented"; return nil }

func (s *_childrenAggregation) Type(relationname string) *_childrenAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_childrenAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_childrenAggregation) ChildrenAggregationCaster() *types.ChildrenAggregation {
	_ = "STUB: not implemented"
	return nil
}
