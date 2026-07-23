package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type _compositeTermsAggregation struct {
	v *types.CompositeTermsAggregation
}

func NewCompositeTermsAggregation() *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) Field(field string) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) MissingBucket(missingbucket bool) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) MissingOrder(missingorder missingorder.MissingOrder) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) Order(order sortorder.SortOrder) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) Script(script types.ScriptVariant) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) ValueType(valuetype valuetype.ValueType) *_compositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeTermsAggregation) CompositeTermsAggregationCaster() *types.CompositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
