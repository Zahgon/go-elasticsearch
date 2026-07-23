package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rangeAggregation struct {
	v *types.RangeAggregation
}

func NewRangeAggregation() *_rangeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_rangeAggregation) Field(field string) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) Format(format string) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) Keyed(keyed bool) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) Missing(missing int) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) Ranges(ranges ...types.AggregationRangeVariant) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) RangesValues(rangesvalues []types.AggregationRange) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) Script(script types.ScriptVariant) *_rangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rangeAggregation) RangeAggregationCaster() *types.RangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
