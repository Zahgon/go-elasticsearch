package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _apiKeyFiltersAggregation struct {
	v *types.ApiKeyFiltersAggregation
}

func NewApiKeyFiltersAggregation() *_apiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) Filters(bucketsapikeyquerycontainer types.BucketsApiKeyQueryContainerVariant) *_apiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) Keyed(keyed bool) *_apiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) OtherBucket(otherbucket bool) *_apiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) OtherBucketKey(otherbucketkey string) *_apiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyFiltersAggregation) ApiKeyFiltersAggregationCaster() *types.ApiKeyFiltersAggregation {
	_ = "STUB: not implemented"
	return nil
}
