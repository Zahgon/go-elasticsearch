package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _compositeAggregation struct {
	v *types.CompositeAggregation
}

func NewCompositeAggregation() *_compositeAggregation { _ = "STUB: not implemented"; return nil }

func (s *_compositeAggregation) After(compositeaggregatekey types.CompositeAggregateKeyVariant) *_compositeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregation) Size(size int) *_compositeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregation) Sources(sources []map[string]types.CompositeAggregationSource) *_compositeAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_compositeAggregation) CompositeAggregationCaster() *types.CompositeAggregation {
	_ = "STUB: not implemented"
	return nil
}
