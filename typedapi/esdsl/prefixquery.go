package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _prefixQuery struct {
	k string
	v *types.PrefixQuery
}

func NewPrefixQuery(field string, value string) *_prefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) CaseInsensitive(caseinsensitive bool) *_prefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) Rewrite(multitermqueryrewrite string) *_prefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) Value(value string) *_prefixQuery { _ = "STUB: not implemented"; return nil }

func (s *_prefixQuery) Boost(boost float32) *_prefixQuery { _ = "STUB: not implemented"; return nil }

func (s *_prefixQuery) QueryName_(queryname_ string) *_prefixQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_prefixQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_prefixQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func NewSinglePrefixQuery() *_prefixQuery { _ = "STUB: not implemented"; return nil }

func (s *_prefixQuery) PrefixQueryCaster() *types.PrefixQuery {
	_ = "STUB: not implemented"
	return nil
}
