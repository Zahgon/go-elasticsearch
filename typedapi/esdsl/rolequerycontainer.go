package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _roleQueryContainer struct {
	v *types.RoleQueryContainer
}

func NewRoleQueryContainer() *_roleQueryContainer { _ = "STUB: not implemented"; return nil }

func (s *_roleQueryContainer) AdditionalRoleQueryContainerProperty(key string, value json.RawMessage) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Bool(bool types.BoolQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Exists(exists types.ExistsQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Ids(ids types.IdsQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Match(key string, value types.MatchQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Range(key string, value types.RangeQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Term(key string, value types.TermQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Terms(terms types.TermsQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_roleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_roleQueryContainer) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}
