package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _wildcardQuery struct {
	k string
	v *types.WildcardQuery
}

func NewWildcardQuery(field string, value string) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) CaseInsensitive(caseinsensitive bool) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) Rewrite(multitermqueryrewrite string) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) Value(value string) *_wildcardQuery { _ = "STUB: not implemented"; return nil }

func (s *_wildcardQuery) Wildcard(wildcard string) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) Boost(boost float32) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) QueryName_(queryname_ string) *_wildcardQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_wildcardQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_wildcardQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleWildcardQuery() *_wildcardQuery { _ = "STUB: not implemented"; return nil }

func (s *_wildcardQuery) WildcardQueryCaster() *types.WildcardQuery {
	_ = "STUB: not implemented"
	return nil
}
