package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _parentAggregation struct {
	v *types.ParentAggregation
}

func NewParentAggregation() *_parentAggregation { _ = "STUB: not implemented"; return nil }

func (s *_parentAggregation) Type(relationname string) *_parentAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_parentAggregation) ParentAggregationCaster() *types.ParentAggregation {
	_ = "STUB: not implemented"
	return nil
}
