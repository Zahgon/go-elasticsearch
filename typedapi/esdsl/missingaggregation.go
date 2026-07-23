package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _missingAggregation struct {
	v *types.MissingAggregation
}

func NewMissingAggregation() *_missingAggregation { _ = "STUB: not implemented"; return nil }

func (s *_missingAggregation) Field(field string) *_missingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_missingAggregation) Missing(missing types.MissingVariant) *_missingAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_missingAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_missingAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_missingAggregation) MissingAggregationCaster() *types.MissingAggregation {
	_ = "STUB: not implemented"
	return nil
}
