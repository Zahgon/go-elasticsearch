package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _userQueryContainer struct {
	v *types.UserQueryContainer
}

func NewUserQueryContainer() *_userQueryContainer { _ = "STUB: not implemented"; return nil }

func (s *_userQueryContainer) AdditionalUserQueryContainerProperty(key string, value json.RawMessage) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Bool(bool types.BoolQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Exists(exists types.ExistsQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Ids(ids types.IdsQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Match(key string, value types.MatchQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) MatchAll(matchall types.MatchAllQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Prefix(key string, value types.PrefixQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Range(key string, value types.RangeQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) SimpleQueryString(simplequerystring types.SimpleQueryStringQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Term(key string, value types.TermQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Terms(terms types.TermsQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) Wildcard(key string, value types.WildcardQueryVariant) *_userQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userQueryContainer) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}
