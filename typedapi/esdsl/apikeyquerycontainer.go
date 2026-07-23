package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _apiKeyQueryContainer struct {
	v *types.ApiKeyQueryContainer
}

func NewApiKeyQueryContainer() *_apiKeyQueryContainer { _ = "STUB: not implemented"; return nil }

func (s *_apiKeyQueryContainer) AdditionalApiKeyQueryContainerProperty(key string, value json.RawMessage) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Bool(bool types.BoolQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Exists(exists types.ExistsQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Ids(ids types.IdsQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Match(key string, value types.MatchQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Range(key string, value types.RangeQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Term(key string, value types.TermQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Terms(terms types.TermsQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_apiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_apiKeyQueryContainer) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}
