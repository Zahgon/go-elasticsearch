package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _apiKeyAggregationContainer struct {
	v *types.ApiKeyAggregationContainer
}

func NewApiKeyAggregationContainer() *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) AdditionalApiKeyAggregationContainerProperty(key string, value json.RawMessage) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Aggregations(aggregations map[string]types.ApiKeyAggregationContainer) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) AddAggregation(key string, value types.ApiKeyAggregationContainerVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Cardinality(cardinality types.CardinalityAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Composite(composite types.CompositeAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) DateRange(daterange types.DateRangeAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Filter(filter types.ApiKeyQueryContainerVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Filters(filters types.ApiKeyFiltersAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Meta(metadata types.MetadataVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Missing(missing types.MissingAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Range(range_ types.RangeAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) Terms(terms types.TermsAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) ValueCount(valuecount types.ValueCountAggregationVariant) *_apiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyAggregationContainer) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}
