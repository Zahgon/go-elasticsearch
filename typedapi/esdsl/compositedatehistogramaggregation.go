package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeDateHistogramAggregation struct {
	v *types.CompositeDateHistogramAggregation
}

func NewCompositeDateHistogramAggregation() *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) CalendarInterval(durationlarge string) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) FixedInterval(durationlarge string) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) Format(format string) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) Offset(duration types.DurationVariant) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) TimeZone(timezone string) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) Field(field string) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) MissingBucket(missingbucket bool) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) Order(order sortorder.SortOrder) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) Script(script types.ScriptVariant) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) ValueType(valuetype valuetype.ValueType) *_compositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeDateHistogramAggregation) CompositeDateHistogramAggregationCaster() *types.CompositeDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
