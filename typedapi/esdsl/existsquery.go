package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _existsQuery struct {
	v *types.ExistsQuery
}

func NewExistsQuery() *_existsQuery { _ = "STUB: not implemented"; return nil }

func (s *_existsQuery) Field(field string) *_existsQuery { _ = "STUB: not implemented"; return nil }

func (s *_existsQuery) Boost(boost float32) *_existsQuery { _ = "STUB: not implemented"; return nil }

func (s *_existsQuery) QueryName_(queryname_ string) *_existsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_existsQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_existsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_existsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_existsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_existsQuery) ExistsQueryCaster() *types.ExistsQuery {
	_ = "STUB: not implemented"
	return nil
}
