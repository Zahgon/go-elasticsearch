package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeHistogramAggregation struct {
	v *types.CompositeHistogramAggregation
}

func NewCompositeHistogramAggregation(interval types.Float64) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) Interval(interval types.Float64) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) Field(field string) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) MissingBucket(missingbucket bool) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) Order(order sortorder.SortOrder) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) Script(script types.ScriptVariant) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) ValueType(valuetype valuetype.ValueType) *_compositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeHistogramAggregation) CompositeHistogramAggregationCaster() *types.CompositeHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
