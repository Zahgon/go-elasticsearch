package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateRangeAggregation struct {
	v *types.DateRangeAggregation
}

func NewDateRangeAggregation() *_dateRangeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_dateRangeAggregation) Field(field string) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) Format(format string) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) Keyed(keyed bool) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) Missing(missing types.MissingVariant) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) Ranges(ranges ...types.DateRangeExpressionVariant) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) RangesValues(rangesvalues []types.DateRangeExpression) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) TimeZone(timezone string) *_dateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateRangeAggregation) DateRangeAggregationCaster() *types.DateRangeAggregation {
	_ = "STUB: not implemented"
	return nil
}
