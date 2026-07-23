package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _reverseNestedAggregation struct {
	v *types.ReverseNestedAggregation
}

func NewReverseNestedAggregation() *_reverseNestedAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reverseNestedAggregation) Path(field string) *_reverseNestedAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reverseNestedAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reverseNestedAggregation) ReverseNestedAggregationCaster() *types.ReverseNestedAggregation {
	_ = "STUB: not implemented"
	return nil
}
